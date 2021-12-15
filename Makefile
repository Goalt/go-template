build:
	docker build -t app:latest .
run:
	docker-compose -p app -f docker-compose.yml up
clean:
	docker stop app ; docker rm app || true
	docker stop app_db_1; docker rm app_db_1 || true
	docker volume rm app_db_data || true
generate:
	cd internal/provider/ && wire ; cd ../..
all: clean build run
