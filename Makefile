.PHONY: run gen

run:
	@go build -o ./tmp/main.exe ./cmd/main.go

gen: 
	@templ generate
