module broker

go 1.22.2

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/net v0.27.0
	shared v0.0.0-00010101000000-000000000000
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/redis/go-redis/v9 v9.6.1 // indirect
	golang.org/x/sys v0.22.0 // indirect
)

replace shared => ../shared
