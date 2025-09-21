# Panduan Setup **microchaos**  
_Project simulasi untuk Chaos Engineering (Docker + Go + MySQL + MongoDB + RabbitMQ + Toxiproxy)_

> Panduan ini menyesuaikan dengan source code & konfigurasi yang sudah ada di repo Anda (`docker-compose.yml`, `.env`, `init.sql`, `seed.sh`, dsb).

---

## 1) Prasyarat

- **Docker** & **Docker Compose v2**
- **Git**
- **Toxiproxy** & **Chaostoolkit** (lihat di [../chaostoolkit](../chaostoolkit/README.md))
- (Opsional) **Go 1.22+** bila ingin menjalankan service langsung

Cek versi:

```bash
docker --version
docker compose version
go version
python --version
```

---

## 2) Struktur Proyek

```
microchaos/
├── .env
├── docker-compose.yml
├── mysql/
│   └── init.sql
├── notifier/
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── order-api/
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── seed.sh
└── README.md
```

### Penjelasan

- **MySQL** → menyimpan data `orders`.
- **MongoDB** → penyimpanan notifikasi hasil konsumsi dari RabbitMQ.
- **RabbitMQ** → event broker (exchange `orders`, queue `order.created`).
- **Toxiproxy** → proxy untuk simulasi chaos (latency, disconnect, dsb).
- **order-api** → API untuk membuat order (insert MySQL, publish event ke RabbitMQ).
- **notifier** → consumer dari RabbitMQ, menyimpan notifikasi ke MongoDB.

---

## 3) Konfigurasi Environment

Semua variable ada di `.env`:

```env
MYSQL_ROOT_PASSWORD=supersecret
MYSQL_DATABASE=app
MYSQL_USER=app
MYSQL_PASSWORD=app

RABBIT_USER=guest
RABBIT_PASS=guest
```

> Network `chaostoolkit_default` sudah otomatis dibuat oleh stack chaostoolkit dan toxiproxy pada step sebelumnya, jadi tidak perlu membuatnya manual.

---

## 4) Jalankan Layanan

1. **Build & start semua container**

```bash
docker compose -f <ROOT_PROJECT>/microchaos/docker-compose.yml up -d 
```

2. **Cek status container**

```bash
docker compose ps
```

3. **Pantau log**

```bash
docker compose logs -f order-api
docker compose logs -f notifier
```

---

## 5) Database

### MySQL (chaosmysql)

Script `mysql/init.sql` otomatis membuat tabel `orders`:

```sql
CREATE TABLE IF NOT EXISTS orders (
  id BIGINT unsigned NOT NULL AUTO_INCREMENT,
  customer VARCHAR(128) NOT NULL,
  item VARCHAR(128) NOT NULL,
  qty INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### MongoDB (chaosmongo)

`notifier` akan menyimpan data notifikasi ke koleksi `notifications` pada database `app`.

---

## 6) Seeding Toxiproxy

Service `toxiproxy-seed` akan menjalankan `seed.sh` untuk membuat proxy:

- MySQL → listen `6061` → upstream `chaosmysql:3306`
- RabbitMQ → listen `6062` → upstream `chaosrabbitmq:5672`
- Mongo → listen `6063` → upstream `chaosmongo:27017`

Log seed bisa dilihat:

```bash
docker logs toxiproxy-seed
```

---

## 7) Akses Layanan

- **Order API** → http://localhost:8081  
- **Notifier API** → http://localhost:8082  
- **RabbitMQ UI** → http://localhost:15672 (user: `guest`, pass: `guest`)

---

## 8) Tes Dasar

### Insert order via Order API

```bash
curl -X POST http://localhost:8081/orders \
  -H "Content-Type: application/json" \
  -d '{"customer":"Alice","item":"Book","qty":2}'
```

- Data order masuk ke MySQL (table `orders`).
- Event `order.created` dikirim ke RabbitMQ.
- Notifier consume event, lalu insert ke MongoDB (`app.notifications`).

### Cek hasil di MySQL

```bash
docker exec -it chaosmysql mysql -uapp -papp -e "SELECT * FROM app.orders\G"
```

### Cek hasil di Mongo

```bash
docker exec -it chaosmongo mongosh app --eval "db.notifications.find().pretty()"
```

---
