up: test docker-start
down: docker-stop

test:
	cd ./orders-management-service && go test ./...
	cd ./users-wallet-service && go test ./...

docker-start:
	docker-compose -f script/docker-compose.yml up -d

docker-stop:
	docker compose -f script/docker-compose.yml stop -t 0