package configs

// type MyType intと宣言するDefined type
// 以前はNamed typeと言っていたが、Go1.11からDefined typeと呼ぶようになった
// type MyType = intと宣言するType alias
// see: https://zenn.dev/bluetree/articles/9d52842dff35cc
// go 1.19から導入された型エイリアス（これで推測され、string()などのキャストが不要になる）
type path = string

const (
	// /api/v1/health
	Health path = "/health"
	// /api/v1/urls
	GetShortURL path = "/urls/:id"
	// /api/v1/urls
	CreateShortURL path = "/urls"
)
