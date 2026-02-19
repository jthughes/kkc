
run:
	@go run .

reset:
	sudo -u postgres dropdb -f kkcbot
	sudo -u postgres psql -c 'create database kkcbot;'
	dbml2sql dbschema.dbml -o ./sql/schema/001_import.sql
	sed -i '1i-- +goose up' ./sql/schema/001_import.sql
	cd ./sql/schema && goose postgres "postgres://postgres:postgres@localhost:5432/kkcbot" up

generate:
	sqlc generate
