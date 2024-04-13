package accesslog

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Echo の middleware.Logger() は、Web アプリケーションのアクセスログを取得して記録するミドルウェアです。このロガーミドルウェアは、リクエストがサーバーに届いた際にその詳細（HTTP メソッド、エンドポイント、リクエスト元の IP アドレス、実行時間、ステータスコードなど）を記録し、開発者がアプリケーションの動作を追跡しやすくする役割を果たします。

// Echo ロガーミドルウェアの主な機能
// リクエスト詳細の記録: 各リクエストについて、リクエストメソッド、URI、リモートアドレス、ユーザーエージェント、ステータスコード、処理時間などの詳細をログとして出力します。
// パフォーマンスモニタリング: リクエストの処理にかかった時間（レイテンシー）を記録することで、アプリケーションのパフォーマンスモニタリングを支援します。
// エラートラッキング: リクエスト処理中に発生したエラーのステータスコードを記録します。
// AccessLog is a middleware to log request and response
func AccessLog() echo.MiddlewareFunc {
	return middleware.Logger()
}
