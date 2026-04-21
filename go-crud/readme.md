❯ go mod init go-crud

> curl -X POST http://localhost:8080/users \
> -H "Content-Type: application/json" \
> -d '{"id": 1, "name": "John Doe", "email": "john@example.com"}'
