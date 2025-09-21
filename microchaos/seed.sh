#!/bin/sh
set -e

echo "🔌 Seeding proxies into Toxiproxy..."

# wait until toxiproxy API is ready
until curl -s http://toxiproxy:8474/proxies >/dev/null 2>&1; do
  echo "⏳ Waiting for Toxiproxy..."
  sleep 2
done

# MySQL proxy
curl -v -X POST -H "Content-Type: application/json" \
  -d '{"name":"mysql_proxy","listen":"0.0.0.0:6061","upstream":"chaosmysql:3306","enabled":true}' \
  http://toxiproxy:8474/proxies || true

# RabbitMQ proxy
curl -v -X POST -H "Content-Type: application/json" \
  -d '{"name":"rabbitmq_proxy","listen":"0.0.0.0:6062","upstream":"chaosrabbitmq:5672","enabled":true}' \
  http://toxiproxy:8474/proxies || true

# Mongo proxy
curl -v -X POST -H "Content-Type: application/json" \
  -d '{"name":"mongo_proxy","listen":"0.0.0.0:6063","upstream":"chaosmongo:27017","enabled":true}' \
  http://toxiproxy:8474/proxies || true

echo "✅ Seeding complete."
