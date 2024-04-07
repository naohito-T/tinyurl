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

lint: sec errcheck staticcheck
	@echo "Check lint finished"

.PHONY: test sec errcheck staticcheck format lint
