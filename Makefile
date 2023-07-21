server:
	go run cmd/main.go

postgres:
	docker run --name perpustakaan -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
	