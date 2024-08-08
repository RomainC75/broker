module consumer

go 1.22.2

require (
	golang.org/x/net v0.27.0
	shared v0.0.0
)

require (
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.22.0 // indirect
)

replace shared => ../shared
