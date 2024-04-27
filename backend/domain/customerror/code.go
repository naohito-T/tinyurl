package customerror

// ApplicationError はアプリケーション特有のエラー情報を保持する基本構造体です。
type ApplicationErrorCode struct {
	// StatusCode（JSON parseでは構造体省略）
	Status int `json:"-"`
	// Code はエラーコードを表します。
	Code string `json:"code"`
	// Message はエラーメッセージを表します。
	Message string `json:"message"`
}

var (
	WrongEmailVerificationCode = ApplicationErrorCode{
		Code:    "WRONG_EMAIL_VERIFICATION_ERROR",
		Message: "Wrong email verification code",
	}

	// Unknown Error: システムがエラーの原因を特定できないか、エラーの種類が分類されていない場合に使用する。
	UnknownCode = ApplicationErrorCode{
		Code:    "unknown_error",
		Message: "unknown_error",
	}

	UnexpectedCode = ApplicationErrorCode{
		Code:    "unexpected_error",
		Message: "unexpected_error",
	}
)
