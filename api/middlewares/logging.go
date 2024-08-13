package middlewares

import (
	"go-myapi/common"
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

/*
http.RequestにはContextが含まれている
ただ、このContextは非公開フィールドであるため、開発者が直接参照する・値をセットすることはできない
= ロギング処理で利用することはできない

net/httpパッケージに含まれていて、Request型に現在セットされているContextを取り出すためのContextメソッドと、
新しくリクエスト型にコンテキストをセットするためのWithContextメソッドが用意されている
*/
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()

		// リクエスト情報をロギング
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := common.SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)
		rlw := NewResLoggingWrighter(w)

		next.ServeHTTP(rlw, req)

		// レスポンス情報をロギング
		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}

// Go公式ではコンテキストが絡んだ処理を行う関数・メソッドには明示的にctxと言うcontext.Context型の第一引数を用意することを推奨している
