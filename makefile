obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

calculator:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

aggregator:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator

up:
	docker-compose up -d

down:
	docker-compose down

.PHONY: obu, aggregator