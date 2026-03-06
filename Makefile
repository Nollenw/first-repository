run: 
	go run main.go
database_up:
	migrate -path migrations -database "postgres://smon:2608@localhost:5432/postgres?sslmode=disable" up
database_down:
	migrate -path migrations -database "postgres://smon:2608@localhost:5432/postgres?sslmode=disable" up
database_up1:
	migrate -path migrations -database "postgres://smon:2608@localhost:5432/postgres?sslmode=disable" up 1