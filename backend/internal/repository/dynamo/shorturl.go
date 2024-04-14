package dynamo

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/naohito-T/tinyurl/backend/domain"
	DynamoClient "github.com/naohito-T/tinyurl/backend/internal/infrastructures/dynamo"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructures/slog"
)

type IShortURLRepository interface {
	Get(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	Put(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type ShortURLRepository struct {
	Client *DynamoClient.Connection
	// インターフェースは既に参照型です。これは、インターフェースが背後でポインタとして機能することを意味し、明示的にポインタとして渡す必要はありません。
	logger slog.ILogger
}

// NewShortURLRepository アスタリスク * は、ポインタを通じてポインタが指し示すアドレスに格納されている実際の値にアクセスするために使います。また、型宣言でポインタ型を示す際にも用いられます。
func NewShortURLRepository(client *DynamoClient.Connection, logger slog.ILogger) *ShortURLRepository {
	// &ShortURLRepository{...} によって ShortURLRepository 型の新しいインスタンスがメモリ上に作成され、そのインスタンスのアドレスが返されます
	return &ShortURLRepository{
		Client: client,
		logger: logger,
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

func (r *ShortURLRepository) GetByShortURL(id string) (domain.ShortURL, error) {
	r.logger.Debug("GetItemInput: %v", id)

	itemKey := ItemKey{
		ID: id,
	}

	av, err := attributevalue.MarshalMap(itemKey)
	if err != nil {
		log.Fatal(err)
		return domain.ShortURL{}, err
	}

	result, err := r.Client.Conn.GetItem(context.TODO(), &dynamodb.GetItemInput{
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

func (r *ShortURLRepository) CreateShortURL(params *domain.ShortURL) error {
	r.logger.Debug("PutItemInput: %v", params)

	item := TableItem{
		ID:          params.ID,
		OriginalURL: params.OriginalURL,
		CreatedAt:   params.CreatedAt,
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatal(err)
	}

	_, err = r.Client.Conn.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("offline-tinyurls"),
		Item:      av,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
