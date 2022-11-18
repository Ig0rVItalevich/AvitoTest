build:
	docker-compose build avito-test

run:
	docker-compose up avito-test

migrate:
	 migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up