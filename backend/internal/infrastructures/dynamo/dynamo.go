package dynamo

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Connection struct {
	Conn *dynamodb.Client
}

func NewDynamoConnection() *Connection {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		slog.Error("unable to load SDK config, %v", err)
	}

	return &Connection{
		Conn: dynamodb.NewFromConfig(cfg),
	}
}
