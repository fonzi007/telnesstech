build:
	go build -o bin/telness-tech-server && \
	chmod +x bin/telness-tech-server

run:
	./bin/telness-tech-server

local-deps-up:
	docker compose -f development/docker-compose/docker-compose.yml up -d --remove-orphans

local-deps-down:
	docker compose -f development/docker-compose/docker-compose.yml down
	docker image rm etc_seed