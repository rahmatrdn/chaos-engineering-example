package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := getenv("SERVICE_PORT", "8082")
	rabbitURL := getenv("RABBIT_URL", "amqp://guest:guest@rabbitmq:5672/")
	queue := getenv("RABBIT_QUEUE", "order.created")
	mongoURI := getenv("MONGO_URI", "mongodb://mongo:27017")
	mongoDB := getenv("MONGO_DB", "app")
	mongoCol := getenv("MONGO_COL", "notifications")

	// Rabbit
	conn, err := amqp.Dial(rabbitURL)
	must(err)
	defer conn.Close()

	ch, err := conn.Channel()
	must(err)
	defer ch.Close()

	_, err = ch.QueueDeclare(queue, true, false, false, false, nil)
	must(err)
	err = ch.QueueBind(queue, "order.created", "orders", false, nil)
	must(err)

	// Mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mc, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	must(err)
	defer mc.Disconnect(context.Background())
	col := mc.Database(mongoDB).Collection(mongoCol)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := mc.Ping(r.Context(), nil); err != nil {
			http.Error(w, "mongo down", 500)
			return
		}
		// Check RabbitMQ connectivity by opening and closing a channel
		testCh, err := conn.Channel()
		if err != nil {
			http.Error(w, "rabbitmq down", 500)
			return
		}
		testCh.Close()
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	go func() {
		log.Println("Queue Start ...")
		msgs, err := ch.Consume(queue, "", true, false, false, false, nil)
		must(err)
		for m := range msgs {
			queueStart := time.Now()
			var any map[string]any
			if err := json.Unmarshal(m.Body, &any); err != nil {
				log.Printf("bad message: %v", err)
				continue
			}
			any["received_at"] = time.Now()
			start := time.Now()
			if _, err := col.InsertOne(context.Background(), any); err != nil {
				log.Printf("mongo insert err: %v", err)
			} else {
				duration := time.Since(start)
				log.Printf("notification stored: %v", any)
				fmt.Printf("Data successfully inserted into MongoDB in %d ms\n", duration.Milliseconds())
			}
			queueDuration := time.Since(queueStart)
			log.Println("Queue processing time:", queueDuration.Milliseconds(), "ms")
		}
	}()

	log.Printf("notifier listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getenv(k, v string) string {
	if s := os.Getenv(k); s != "" {
		return s
	}
	return v
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
