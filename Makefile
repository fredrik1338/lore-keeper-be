run:
	go run main.go

build:
	go build -o bin/main main.go

docker-build-dev:
	docker build -t lore-keeper-be-dev .

docker-build-prod:
	docker build -t lore-keeper-be-prod .

docker-run-dev:
	docker run -p 8080:8080 lore-keeper-be-dev

define generate_timestamp
$(shell date +'%Y%m%d%H%M')
endef

docker-tag-and-push-dev:
	docker tag lore-keeper-be-dev ghcr.io/fredrik1338/lore-keeper-be-dev:${generate_timestamp}
	docker push ghcr.io/fredrik1338/lore-keeper-be-dev:${generate_timestamp}
	docker tag lore-keeper-be-dev ghcr.io/fredrik1338/lore-keeper-be-dev:latest
	docker push ghcr.io/fredrik1338/lore-keeper-be-dev:latest