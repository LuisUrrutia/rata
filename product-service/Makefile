protoc:
	protoc -I. --micro_out=. --go_out=. proto/product/product.proto

run-bin:
	DB_HOST="127.0.0.1" DB_USER="root" DB_NAME="scraper" DB_PASSWORD="secret" ./product-service --broker=nats --registry=mdns