postgres:
	docker run --name postgres16 -e POSTGRES_USER=root_user -e POSTGRES_PASSWORD=123456 -p 5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root_user --owner=root_user gobankhub

dropdb:
	docker exec -it postgres16 dropdb gobankhub

migrateup:
	migrate -path db/migration -database "postgresql://root_user:123456@localhost:5432/gobankhub?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root_user:123456@localhost:5432/gobankhub?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: createdb dropdb postgres migrateup migratedown sqlc