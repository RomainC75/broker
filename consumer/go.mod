module consumer

go 1.22.2

require (
	golang.org/x/net v0.27.0
	shared v0.0.0-00010101000000-000000000000
)

require github.com/joho/godotenv v1.5.1 // indirect

replace shared => ../shared
