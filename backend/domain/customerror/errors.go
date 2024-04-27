package customerror

import (
	"fmt"
	"net/http"
	"net/url"
)

// ApplicationError はアプリケーション特有のエラー情報を保持する基本構造体です。
type ApplicationError struct {
	// StatusCode（JSON parseでは構造体省略）
	Status int `json:"-"`
	// Code はエラーコードを表します。
	Code string `json:"code"`
	// Message はエラーメッセージを表します。
	Message string `json:"message"`
}

// Error メソッドは error インターフェースを実装します。
func (e *ApplicationError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}

// WrongEmailVerificationError は特定の条件下で利用されるカスタムエラーです。
type WrongEmailVerificationError struct {
	ApplicationError
	IsTally       bool
	IsRedirect    bool
	RedirectQuery *url.Values
}

// NewWrongEmailVerificationError は WrongEmailVerificationError を生成します。
func NewWrongEmailVerificationError(isTally, isRedirect bool, redirectQuery string) *WrongEmailVerificationError {
	q, _ := url.ParseQuery(redirectQuery)
	return &WrongEmailVerificationError{
		ApplicationError: ApplicationError{
			Status:  http.StatusBadRequest,
			Code:    "WRONG_EMAIL_VERIFICATION_ERROR",
			Message: "Wrong email verification code",
		},
		IsTally:       isTally,
		IsRedirect:    isRedirect,
		RedirectQuery: &q,
	}
}

// LogError はエラーのロギングを行うメソッドです。
// func (e *WrongEmailVerificationError) LogError(req *http.Request) {
// 	fmt.Println("Error occurred:", e.Error(), "Query:", e.RedirectQuery.Encode())
// }

// Is メソッドを追加
// func (e *WrongEmailVerificationError) Is(target error) bool {
// 	t, ok := target.(*WrongEmailVerificationError)
// 	return ok && e.Code == t.Code
// }

// WrongEmailVerificationErrorInstance は WrongEmailVerificationError のグローバルインスタンスです。
var WrongEmailVerificationErrorInstance = NewWrongEmailVerificationError(false, false, "")
