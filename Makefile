dev:
	air
build:
	go build -o ./dist/doctors ./cmd/doctors/main.go
create-migration:
	migrate create -ext sql -dir ./pkg/db/migrations -seq $(name)
revert-migration:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down