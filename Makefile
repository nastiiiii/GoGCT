createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:beetroot@dB:5432/GCT-SYSTEM?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:beetroot@dB:5432/GCT-SYSTEM?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

hello:
	echo "Hello"

createDocker:
	docker run -d --name GCTproject -e POSTGRES_USER=root -e POSTGRES_PASSWORD=beetroot -e POSTGRES_DB=GCT -p 5433:5432 postgres

copyDB:
	cat /Users/nasaska/Downloads/GCT-SYSTEM.sql | docker exec -i GCTproject psql -U root -d GCT

.PHONY: createdb createdb dropdb migratedown migrateup sqlc createDocker copyDB