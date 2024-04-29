package context

// package humaecho

// import (
// 	"context"

// 	"github.com/danielgtaylor/huma/v2"
// 	"github.com/labstack/echo/v4"
// )

// // EchoHumaContext は huma.Context に echo.Context を追加します。
// type EchoHumaContext interface {
// 	// これは継承らしい
// 	huma.Context
// 	EchoContext() echo.Context
// }

// // echoContextImpl は EchoHumaContext の実装です。
// type echoContextImpl struct {
// 	humaCtx huma.Context
// 	echoCtx echo.Context
// }

// func (e *echoContextImpl) Operation() *huma.Operation {
// 	return e.humaCtx.Operation()
// }

// func (e *echoContextImpl) Context() context.Context {
// 	return e.humaCtx.Context()
// }

// // 以下、huma.Context の他のメソッドも同様に実装します。
// // ...

// // EchoContext は echo.Context を返します。
// func (e *echoContextImpl) EchoContext() echo.Context {
// 	return e.echoCtx
// }

// // NewEchoHumaContext は新しい EchoHumaContext を作成します。
// func NewEchoHumaContext(hCtx huma.Context, eCtx echo.Context) EchoHumaContext {
// 	return &echoContextImpl{
// 		humaCtx: hCtx, // huma.Context をセット
// 		echoCtx: eCtx, // echo.Context をセット
// 	}
// }
