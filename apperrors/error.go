package apperrors

type MyAppError struct {
	ErrCode        // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ

	// Error() stringメソッドを持つ構造体はerrorインターフェースを満たし、独自エラー型として扱うことができる
	Err error `json:"-"` // エラーチェーンのための内部エラー `json:"-"`:jsonエンコードされないように指定
}

// Errorメソッドは、その構造体をエラーとして扱うためにMUSTでつけるメソッド
// 本体のErrorメソッドの役割は、
// 「そのエラーがfmt.Print系関数等で出力されるときにどのような文字列になるか」ということを決めるためのもの
func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
