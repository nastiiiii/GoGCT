test:
	go test -v -cover ./...

hello:
	echo "Hello"

createDocker:
	docker run -d --name GCTproject -e POSTGRES_USER=root -e POSTGRES_PASSWORD=beetroot -e POSTGRES_DB=GCT -p 5433:5432 postgres

copyDB:
	cat /Users/nasaska/Desktop/GCT-SYSTEM.sql | docker exec -i GCTproject psql -U root -d GCT

.PHONY: createdb createdb dropdb migratedown migrateup sqlc createDocker copyDB