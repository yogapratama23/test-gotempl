run:
	@templ generate
	@go run cmd/main.go

generate:
	@templ generate

watch:
	@gow run cmd/main.go