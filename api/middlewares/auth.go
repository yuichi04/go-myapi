package middlewares

import (
	"errors"
	"go-myapi/apperrors"
	"go-myapi/common"
	"os"
	"strings"

	"google.golang.org/api/idtoken"

	"context"
	"net/http"
)

var googleClientID = os.Getenv("GOOGLE_CLIENT_ID")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// ヘッダからAuthorizationフィールドを抜き出す
		authorizatioin := req.Header.Get("Authorization")

		// Authorization フィールドが"Bearer [IDトークン]"の形になっているか検証
		authHeaders := strings.Split(authorizatioin, " ") // 空白区切り文字で２つに分かれているか
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthrizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1] // 空白区切りで分けた１つ目がBearerで、２つ目が空でないか
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthrizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// IDトークン検証
		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err = apperrors.Unauthrizated.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err = apperrors.Unauthrizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// name フィールドを payload から抜き出す
		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthrizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		req = common.SetUserName(req, name.(string))

		next.ServeHTTP(w, req)
	})
}
