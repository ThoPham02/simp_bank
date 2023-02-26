postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root simp_bank
dropdb:
	docker exec -it postgres dropdb simp_bank
migrate:
	migrate create -ext sql -dir db/migration -seq add_users
migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simp_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simp_bank?sslmode=disable" -verbose down
sqlcinit:
	sudo docker run --rm -v "$$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc init
sqlc:
	sudo docker run --rm -v "$$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate
test:
	go test -v -cover ./...
mock:
	mockgen -destination ./db/mock/store.go -package mockdb github.com/ThoPham02/simp_bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc sqlcinit test mock