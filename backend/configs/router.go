package configs

// type MyType intと宣言するDefined type
// 以前はNamed typeと言っていたが、Go1.11からDefined typeと呼ぶようになった
// type MyType = intと宣言するType alias
// see: https://zenn.dev/bluetree/articles/9d52842dff35cc
// go 1.19から導入された型エイリアス（これで推測され、string()などのキャストが不要になる）
type path = string

const (
	Health         path = "/health"
	GetShortURL    path = "/api/v1/urls/:shortUrl"
	CreateShortURL path = "/api/v1/urls"
)
