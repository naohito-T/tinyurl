package customerror

// https://zenn.dev/msksgm/articles/20220325-unwrap-errors-is-as

type ValidationError struct {
	Message string
	Err     error
}

func (ve *ValidationError) Error() string {
	return ve.Message
}

func (ve *ValidationError) Unwrap() error {
	return ve.Err
}
