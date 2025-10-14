run:
	@echo "Executando o projeto..."
	go run main.go

ti:
	@echo "Executando mod tidy"
	go mod tidy

fm:
	@echo "formatando codigo"
	go fmt
