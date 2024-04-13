package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	DynamoClient "github.com/naohito-T/tinyurl/backend/internal/infrastructures/dynamo"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructures/slog"
)

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

func (r *ShortURLRepository) Put(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	r.logger.Debug("PutItemInput: %v", params)
	return r.Client.Conn.PutItem(ctx, params, optFns...)
}

func (r *ShortURLRepository) Get(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	r.logger.Debug("GetItemInput: %v", params)
	return r.Client.Conn.GetItem(ctx, params, optFns...)
}
