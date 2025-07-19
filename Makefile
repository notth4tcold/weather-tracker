APP_NAME=weather-tracker
PORT=8080
CONTAINER_NAME=$(APP_NAME)-dev
COMPOSE_FILE=docker/docker-compose.yml

# build:
# 	docker build -t $(APP_NAME) .

# run:
# 	docker run -d --name $(CONTAINER_NAME) -p $(PORT):$(PORT) $(APP_NAME)

# stop:
# 	docker stop $(CONTAINER_NAME) || true
# 	docker rm $(CONTAINER_NAME) || true

# logs:
# 	docker logs -f $(CONTAINER_NAME)

# clean:
# 	docker rm $(CONTAINER_NAME) || true
# 	docker rmi $(APP_NAME) || true

# reload: stop clean build run

build:
	docker-compose -f $(COMPOSE_FILE) build

up:
	docker-compose -f $(COMPOSE_FILE) up -d

down:
	docker-compose -f $(COMPOSE_FILE) down

clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans
	docker rm -f weather-pgadmin weather-db weather-app || true
	docker rmi $(APP_NAME) || true

migrate-up:
	migrate -path migrations -database "postgres://mei:mei@localhost:5432/weatherdb?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://mei:mei@localhost:5432/weatherdb?sslmode=disable" down

ps: 
	docker-compose -f $(COMPOSE_FILE) ps

reload: down clean build up