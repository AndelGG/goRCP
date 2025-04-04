run:
	go run cmd/sso/main.go --config=./config/local.yaml
migrate:
	go run ./cmd/migrator/main.go --storage=./storage/sso.db --migrations=./migrations
migTest:
	go run ./cmd/migrator/main.go --storage=./storage/sso.db --migrations=./tests/migrations --table=migrations_test