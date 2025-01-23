.PHONY: generate

run:
	go run ./main/main.go

dev_containers:
	docker-compose down
	docker-compose up -d

swagger:
	swag init -g ./main/main.go

generate:
	go generate ./generate/main