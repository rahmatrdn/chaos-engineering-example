package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	Customer string `json:"customer"`
	Item     string `json:"item"`
	Qty      int    `json:"qty"`
}

func main() {
	port := getenv("SERVICE_PORT", "8081")
	mysqlDSN := getenv("MYSQL_DSN", "app:app@tcp(mysql:3306)/app?parseTime=true")
	rabbitURL := getenv("RABBIT_URL", "amqp://guest:guest@rabbitmq:5672/")
	exchange := getenv("RABBIT_EXCHANGE", "orders")

	fmt.Println("Service start")

	db, err := sql.Open("mysql", mysqlDSN)
	must(err)
	defer db.Close()

	// RabbitMQ
	conn, err := amqp.Dial(rabbitURL)
	must(err)
	defer conn.Close()

	ch, err := conn.Channel()
	must(err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-delete
		false,    // internal
		false,    // no-wait
		nil,
	)
	must(err)

	// Routes
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(w, "db down", 500)
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

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request create order start ...")
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", 405)
			return
		}
		var o Order
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			http.Error(w, "bad json", 400)
			return
		}
		startReq := time.Now()
		res, err := db.Exec("INSERT INTO orders (customer, item, qty, created_at) VALUES (?,?,?,?)",
			o.Customer, o.Item, o.Qty, time.Now())
		if err != nil {
			http.Error(w, "db insert err: "+err.Error(), 500)
			return
		}

		id, _ := res.LastInsertId()
		body, _ := json.Marshal(map[string]any{
			"id":       id,
			"customer": o.Customer,
			"item":     o.Item,
			"qty":      o.Qty,
		})
		startPublish := time.Now()
		if err := ch.Publish(exchange, "order.created",
			false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			}); err != nil {
			http.Error(w, "publish err: "+err.Error(), 500)
			return
		}

		publishDuration := time.Since(startPublish)
		totalDuration := time.Since(startReq)
		fmt.Printf("Successfully sent to RabbitMQ order.created in %vms\n", publishDuration.Milliseconds())
		log.Printf("Create order request completed in %dms\n", totalDuration.Milliseconds())

		w.WriteHeader(201)
		w.Write([]byte(`{"status":"created","id":` + itoa(id) + `}`))
	})

	log.Printf("order-api listening on :%s", port)
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

func itoa(id int64) string {
	return json.Number((json.Number)(string([]byte(fmtInt(id))))).String()
}

func fmtInt(i int64) []byte {
	// tiny alloc-free int64 -> []byte
	if i == 0 {
		return []byte("0")
	}
	var b [20]byte
	pos := len(b)
	n := i
	if n < 0 {
		n = -n
	}
	for n > 0 {
		pos--
		b[pos] = byte('0' + n%10)
		n /= 10
	}
	if i < 0 {
		pos--
		b[pos] = '-'
	}
	return b[pos:]
}
