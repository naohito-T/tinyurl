FROM golang:1.21.1

WORKDIR /app

RUN apt update -yqq \
    && go install github.com/cosmtrek/air@v1.29.0

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app/

RUN go build -o ./main

CMD ["air", "-c", ".air.toml"]
