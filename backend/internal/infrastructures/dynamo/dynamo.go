package dynamo

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Connection struct {
	Conn *dynamodb.Client
}

func NewDynamoConnection() *Connection {
	// https://zenn.dev/y16ra/articles/40ff14e8d2a4db
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           "http://aws:4566", // LocalStackのDynamoDBエンドポイント
			SigningRegion: "ap-northeast-1",
		}, nil
	})
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-1"),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		slog.Error("unable to load SDK config, %v", err)
	}

	return &Connection{
		Conn: dynamodb.NewFromConfig(cfg),
	}
}
