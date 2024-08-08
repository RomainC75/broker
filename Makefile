run:
	docker compose up

stop:
	docker compose down

migrateup:
	migrate --path shared/db/migration --database "postgresql://admin:azerty@localhost:5432/cryptos?sslmode=disable" --verbose up

migratedown:
	migrate --path shared/db/migration --database "postgresql://admin:azerty@localhost:5432/cryptos?sslmode=disable" --verbose down

sqlc:
	cd server && sqlc generate

test:
	cd server && go test -v -cover ./...

server:
	cd server && go run main.go

mock:
	cd server && mockgen -package mockdb -destination db/mock/store.go github.com/RomainC75/server-garbage/db/sqlc Store

.PHONY: run stop migrateup migratedown sqlc test server