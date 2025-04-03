
APP_NAME := mini-billing
DOCKER_IMAGE := $(APP_NAME):latest
CONTAINER_NAME := $(APP_NAME)-container
PORT := 8181

build:
	-go build -o ./tmp/main cmd/main.go
dev:
	# - lsof -ti :8080 | xargs kill -9
	- make init
	- air -c .air.toml
run:
	go run cmd/main.go


docker:
	@echo "🔄 Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

	@echo "🛑 Stopping old container (if exists)..."
	- docker stop $(CONTAINER_NAME) 2>/dev/null || true
	- docker rm $(CONTAINER_NAME) 2>/dev/null || true

	@echo "🚀 Running new container..."
	docker run -d --name $(CONTAINER_NAME) -p $(PORT):$(PORT) $(DOCKER_IMAGE)

	@echo "✅ Done! App is running at http://localhost:$(PORT)"

logs:
	docker logs -f $(CONTAINER_NAME)

stop:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true
