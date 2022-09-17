createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	./sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	./mockgen.exe -package mockdb -destination db/mock/store.go github.com/Elessurgon/go_pract/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown sqlc test server mock