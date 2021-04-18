up:
	docker-compose up -d --scale go-orm=3
build:
	docker build -t go-orm .
	docker build -f Dockerfile.nginx -t go-orm-nginx .
down:
	docker-compose down
ps:
	docker-compose ps