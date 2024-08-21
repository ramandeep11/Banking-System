postgres:
	docker run --name postgress12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine 

createdb:
	docker exec -it postgress12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgress12 dropdb simple_bank

migrateup:
	migrate -verbose -path ./db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" up

migrateup1:
	migrate -verbose -path ./db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" up 1

migratedown:
	migrate -path ./db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path ./db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test ./db/sqlc/ -v -cover

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc server mock