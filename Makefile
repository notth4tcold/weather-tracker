APP_NAME=weather-tracker
PORT=8080
CONTAINER_NAME=$(APP_NAME)-dev

build:
	docker build -t $(APP_NAME) .

run:
	docker run -d --name $(CONTAINER_NAME) -p $(PORT):$(PORT) $(APP_NAME)

stop:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

logs:
	docker logs -f $(CONTAINER_NAME)

clean:
	docker rm $(CONTAINER_NAME) || true
	docker rmi $(APP_NAME) || true

reload: stop clean build run
