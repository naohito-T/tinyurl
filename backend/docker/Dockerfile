FROM golang:1.21.1

WORKDIR /app

RUN apt update -yqq \
    && go install github.com/cosmtrek/air@v1.29.0

COPY go.mod go.sum /app/
# 依存関係をダウンロード
RUN go mod download
COPY . /app/
# アプリケーションをビルド
RUN go build -o ./main
# 実行コマンドを設定
CMD ["air", "-c", ".air.toml"]
