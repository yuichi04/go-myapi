package middlewares

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWrighter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// ハンドラがHTTPレスポンスコードを書き込むときに使うメソッド
func (rsw *resLoggingWriter) WriteHeader(code int) {
	// resLoggingWriter構造体のcodeフィールドに使うレスポンスコードを保存する
	rsw.code = code

	// HTTPレスポンスに使うレスポンスコードを指定
	// (=WriteHeaderメソッド本来の機能を呼び出し)
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// リクエスト情報をロギング
		log.Println("req: ", req.RequestURI, req.Method)

		rlw := NewResLoggingWrighter(w)

		next.ServeHTTP(rlw, req)

		log.Println("res: ", rlw.code)
	})
}
