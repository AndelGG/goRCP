run:
	go run cmd/sso/main.go --config=./config/local.yaml
migrate:
	go run ./cmd/migrator --storage=./storage/sso.db --migrations=./migrations