build:
	@cd server && \
		echo "---------------------------------------------" && \
		echo "Building server..." && \
		go mod tidy && \
		go build -o ./.bin/main main.go && \
		echo "Server build complete." && \
		echo "---------------------------------------------"
	@echo "Binary location: $(shell pwd)/server/.bin/main"
	@echo "---------------------------------------------"

run:
	@echo "---------------------------------------------"
	@cd server && \
		echo "Running server..." && \
		go run main.go
	@echo "Server is running."
	@echo "---------------------------------------------"

docker-build:
	@echo "---------------------------------------------"
	@echo "Building docker image..."
	@cd server && \
		echo "Building Docker image..." && \
		sudo docker build -t edd-app:latest . && \
		echo "Docker image built."
	@echo "Docker image hash: $(shell sudo docker images -q edd-app:latest)"
	@echo "---------------------------------------------"