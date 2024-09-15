package infrastructure

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/naohito-T/tinyurl/backend/configs"
)

type Client interface {
	Get(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	Put(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	Search(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}

type Connection struct {
	*dynamodb.Client
	ILogger
	configs.AppEnvironment
}

func NewDynamoConnection(logger ILogger, env configs.AppEnvironment) *Connection {
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
		logger.Error("unable to load SDK config, %v", err)
	}

	return &Connection{
		dynamodb.NewFromConfig(cfg),
		logger,
		env,
	}
}

func (c *Connection) Get(ctx context.Context, params *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	c.Info("GetItemInput: %v", params)
	if result, err := c.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(c.GetTinyURLCollectionName()),
		Key:       params.Key,
	}); err != nil {
		c.Error("Get error: %v", err)
		return nil, err
	} else {
		return result, nil
	}
}

func (c *Connection) Put(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	c.Info("PutItemInput: %v", params)
	if result, err := c.PutItem(ctx, &dynamodb.PutItemInput{
		// rateLimitなどテーブルnameは上から渡したほうがいいかも
		TableName: aws.String(c.GetTinyURLCollectionName()),
		Item:      params.Item,
	}, optFns...); err != nil {
		c.Error("Put error: %v", err)
		return nil, err
	} else {
		return result, nil
	}
}

// // QueryInput の構築
//
//	input := &dynamodb.QueryInput{
//		TableName:              aws.String("offline-tinyurls"),
//		IndexName:              aws.String("OriginalURLIndex"), // OriginalURL に基づいた GSI (グローバルセカンダリインデックス) の名前
//		KeyConditionExpression: aws.String("OriginalURL = :originalURL"),
//		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
//			":originalURL": {S: aws.String(originalURL)},
//		},
//	}
func (c *Connection) Search(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	c.Info("SearchItemInput: %v", params)
	if result, err := c.Query(ctx, params, optFns...); err != nil {
		c.Error("Query error: %v", err)
		return nil, err
	} else {
		return result, nil
	}
}
