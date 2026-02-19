all: reset run

run:
	@go run .

reset:
	sudo -u postgres dropdb -f kkc
	sudo -u postgres psql -c 'create database kkc;'
	dbml2sql dbschema.dbml -o ./sql/schema/001_import.sql
	sed -i '1i-- +goose up' ./sql/schema/001_import.sql
	cd ./sql/schema && goose postgres "postgres://postgres:postgres@localhost:5432/kkc" up

generate:
	sqlc generate
