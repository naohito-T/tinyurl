package dynamo

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/naohito-T/tinyurl/backend/domain"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructure"
)

type IShortURLRepository interface {
	Get(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	Put(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type ShortURLRepository struct {
	*infrastructure.Connection
	// インターフェースは既に参照型です。これは、インターフェースが背後でポインタとして機能することを意味し、明示的にポインタとして渡す必要はありません。
	infrastructure.ILabelLogger
}

func NewShortURLRepository(client *infrastructure.Connection, logger infrastructure.ILabelLogger) *ShortURLRepository {
	// &ShortURLRepository{...} によって ShortURLRepository 型の新しいインスタンスがメモリ上に作成され、そのインスタンスのアドレスが返されます
	return &ShortURLRepository{
		client,
		logger,
	}
}

type TableItem struct {
	ID          string `dynamodbav:"id"`
	OriginalURL string `dynamodbav:"originalURL"`
	CreatedAt   string `dynamodbav:"createdAt"`
}

type ItemKey struct {
	ID string `dynamodbav:"id"`
}

// 構造体に属することで、構造体が初期されていないと呼び出すことはできないことになる。
func (r *ShortURLRepository) Get(ctx context.Context, hashURL string) (domain.ShortURL, error) {
	// r.Debug("GetItemInput: %v", hashURL)

	itemKey := ItemKey{
		ID: hashURL,
	}

	av, err := attributevalue.MarshalMap(itemKey)
	if err != nil {
		log.Fatal(err)
		return domain.ShortURL{}, err
	}

	result, err := r.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("offline-tinyurls"),
		Key:       av,
	})

	if err != nil {
		log.Fatal(err)
		return domain.ShortURL{}, err
	}

	var item TableItem
	err = attributevalue.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Fatal(err)
		return domain.ShortURL{}, err
	}

	shortURL := domain.ShortURL{
		ID:          item.ID,
		OriginalURL: item.OriginalURL,
		CreatedAt:   item.CreatedAt,
	}

	return shortURL, nil
}

func (r *ShortURLRepository) Put(ctx context.Context, params *domain.ShortURL) (domain.ShortURL, error) {
	// r.Info("PutItemInput: %v", params)

	item := TableItem{
		ID:          params.ID,
		OriginalURL: params.OriginalURL,
		CreatedAt:   params.CreatedAt,
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatal(err)
		return domain.ShortURL{}, err // エラー時にゼロ値を返す
	}

	_, err = r.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("offline-tinyurls"),
		Item:      av,
	})
	var oe *smithy.OperationError
	if errors.As(err, &oe) {
		log.Printf("failed to call service: %s, operation: %s, error: %v", oe.Service(), oe.Operation(), oe.Unwrap())
	}
	if err != nil {
		log.Fatal(err)
		return domain.ShortURL{}, err // エラー時にゼロ値を返す
	}
	shortURL := domain.ShortURL{
		ID:          params.ID,
		OriginalURL: params.OriginalURL,
		CreatedAt:   params.CreatedAt,
	}
	return shortURL, nil
}
