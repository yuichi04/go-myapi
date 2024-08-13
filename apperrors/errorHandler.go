package apperrors

import (
	"encoding/json"
	"errors"
	"go-myapi/common"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	// エラーの種類を判別して、適切なhttpレスポンスを返す
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := common.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	case RequiredAuthrizationHeader, Unauthrizated:
		statusCode = http.StatusUnauthorized
	case NotMatchUser:
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	// WriteHeader: リクエストヘッダにステータスコードを書き込むメソッド
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
