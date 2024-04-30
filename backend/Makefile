# .env 関連
STAGE ?= local

dev:
	@echo "Start develop"
	air -c .air.toml
	@echo "Start localstack finished"

# go標準ライブラリのテスト（vetも実行する）
test:
	go test ./... -v

# see https://logmi.jp/tech/articles/324549
# セキュリティチェック
sec:
	golangci-lint run --enable-only=gosec ./...

# エラーチェック
errcheck:
	golangci-lint run --enable-only=errcheck ./...

# 静的解析
staticcheck:
	golangci-lint run --enable-only=staticcheck ./...

# フォーマット
format:
	golangci-lint run --enable-only=goimports --fix ./...

lint:
	golangci-lint run ./...
	@echo "Check lint finished"

lint-f:
	golangci-lint run --fix ./...
	@echo "Check lint with Write finished"

build:
	GOARCH=arm64 GOOS=linux go build -o bin/main cmd/api/main.go

build-openapi:
	go run ./cmd/api/main.go openapi >./api/openapi.yaml

# --------------------------
# Utils
# --------------------------

# DynamoDBの管理画面を起動
dynamo.admin:
	export DYNAMO_ENDPOINT="http://localhost:54566" && \
	export PORT="4005" && \
	npx dynamodb-admin

.PHONY: dev test sec errcheck staticcheck format lint-all lint build build-openapi dynamo.admin
