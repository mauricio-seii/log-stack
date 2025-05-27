APP_NAME=logparser
LOGS_DIR=./logs
PARSED_DIR=./logs/parsed
DOCKER_COMPOSE_FILE=docker-compose.yml

.PHONY: all build up clean logs

all: build up

build:
	@echo "🛠️  Building Go application Docker image..."
	docker build -t $(APP_NAME) ./logparser

prepare:
	@echo "📁 Creating required directories..."
	mkdir -p $(LOGS_DIR)
	mkdir -p $(PARSED_DIR)

up: prepare
	@echo "🚀 Starting Docker Compose environment..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

logs:
	@echo "📜 Tailing logs from logparser..."
	docker-compose logs -f $(APP_NAME)

clean:
	@echo "🧹 Stopping and removing containers, networks, volumes..."
	docker-compose down --volumes
	rm -rf $(LOGS_DIR) $(PARSED_DIR)
