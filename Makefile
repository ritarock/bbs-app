MOCK_GEN = go run -mod=mod go.uber.org/mock/mockgen@latest

mock:
	$(MOCK_GEN) \
	-source=./domain/repository/post_repository.go \
	-destination=./testing/mock/post_repository.go \
	-package=mock

test:
	go test ./...

create.table:
	sqlite3 bbs.db < ./sqlc/schema.sql

sqlc.pull:
	docker pull sqlc/sqlc

sqlc.generate:
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc generate

tsp.build:
	docker compose build

tsp.init:
	docker compose run --rm typespec tsp init

tsp.compile:
	docker compose run --rm typespec tsp compile .

ogen:
	go run -mod=mod github.com/ogen-go/ogen/cmd/ogen@latest \
	-package api \
	-target infra/handler/api \
	-clean ./tsp-output/schema/openapi.yaml
