module broker

go 1.22.2

require (
	github.com/joho/godotenv v1.5.1
	golang.org/x/net v0.27.0
	shared v0.0.0-00010101000000-000000000000
)

replace shared => ../shared
