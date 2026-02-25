build:
	go build -o openslides-backend-service .

run:
	go run . run

run-tests:
	go test ./...

generate-models:
	go run ./cmd/gen/ -models -path ../lib/openslides-go/meta/models.yml > internal/model/generated.go

generate-schemas:
	go run ./cmd/gen/ -schemas -path ../lib/openslides-go/meta/models.yml > internal/model/schema_generated.go

generate: generate-models generate-schemas

fmt:
	go fmt ./...

vet:
	go vet ./...

.PHONY: build run run-tests generate-models generate-schemas generate fmt vet
