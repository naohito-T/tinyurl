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
	// リダイレクト用エンドポイント
	GetShortURL path = "/urls/:id"
	// 短縮URLを作成するためのエンドポイント
	CreateShortURL path = "/urls"
	// すべてのURLをリストするためのエンドポイント
	ListShortURLs path = "/urls/list"
	//  特定の短縮URLの詳細情報を取得
	GetOnlyShortURL path = "/urls/info/:id"
)
