# .env 関連
STAGE ?= local

# go標準ライブラリのテスト（vetも実行する）
test:
	go test ./... -v

# see https://logmi.jp/tech/articles/324549
# セキュリティチェック
sec:
	gosec ./...

# エラーチェック
errcheck:
	errcheck ./...

# 静的解析
staticcheck:
	staticcheck ./...

# フォーマットとインポートの整理
format:
	go fmt ./...
	goimports -w ./

# さまざまなlintを包括したサードパーティツール
lint-all:
	golangci-lint run ./...

lint: sec errcheck staticcheck
	@echo "Check lint finished"

build:
	GOARCH=arm64 GOOS=linux go build -o bin/main cmd/api/main.go

# --------------------------
# Utils
# --------------------------

# DynamoDBの管理画面を起動
dynamo.admin:
	export DYNAMO_ENDPOINT="http://localhost:54566" && \
	export PORT="4005" && \
	npx dynamodb-admin

.PHONY: test sec errcheck staticcheck format lint-all lint build test dynamo.admin
