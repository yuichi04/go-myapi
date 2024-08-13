package main

import (
	"database/sql"
	"fmt"
	"go-myapi/api"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	// 1. サーバー全体で使用するsql.DB型を1つ生成する
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	// 2. コントローラ型MyAppControllerのハンドラメソッドとパストの関連付けを行う
	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// package main

// import (
// 	"crypto"
// 	"crypto/rsa"
// 	"crypto/sha256"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"math/big"
// 	"strings"
// )

// func main() {
// 	// IDトークンから、ヘッダー・ペイロードを入手するプログラム

// 	idToken := `eyJhbGciOiJSUzI1NiIsImtpZCI6ImQyZDQ0NGNmOGM1ZTNhZTgzODZkNjZhMTNhMzE2OTc2YWEzNjk5OTEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIxMDIwNzQ5NjQ0OTg1LTJyMGEwZmg2MGhwODcxdDdnanBycXF1NHZjMjBlczFmLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiMTAyMDc0OTY0NDk4NS0ycjBhMGZoNjBocDg3MXQ3Z2pwcnFxdTR2YzIwZXMxZi5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjEwNzI4NDIzNTU3MjY1OTEyNTg3NyIsIm5vbmNlIjoiMTExMTEiLCJuYmYiOjE3MjM1MTQxNzUsIm5hbWUiOiJZIFkiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jSmNvRlpWZklFSUVrZXY2bk9qbmlCU2VrREpwRmNHUlQ5WmR5RHBiUDVsVkw3dkRnPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6IlkiLCJmYW1pbHlfbmFtZSI6IlkiLCJpYXQiOjE3MjM1MTQ0NzUsImV4cCI6MTcyMzUxODA3NSwianRpIjoiODY4MGNkMjliZGYxYjQ4YTA3MjcwYzI1MGY5MjUyZTU2NjM5OTdlMyJ9.LHak307VUrClJrv3fPdYE-aFsLDGhBAwvxoD5onIny6bLPd17Onbqpyh0PQlOms26ZgeEC9KDmBzBmIY6iHXwUiUM87Qy0g1Xv0f6-KQxkBdUciahDksGb6_Lwggj31e7q6FJ03gV79-xOzq53Ho_D46pxFtnrwVeRHzGRM-NqixDZrjD5rFmPuZiJWkfAkAb2kPo1m0qdusyI_BU3i_nLlUu_KnIna7RQthhxnRxOTs9sNya7G-q75K0EzaP2yerkizGq_pFtMce4_brsd78wWSuiwVo7hexl4K4gqhINy2XOUwheYEdZ_W_yvMtv26JrHvrjVpWS6D-0Qpw-_nwQ&authuser=0&prompt=consent&version_info=CmxfU1ZJX0VQN2g0T0R2OEljREdBNGlQMDFCUlVSSVpsOXhZa2wyYW5sVFYzVlplVWRJU1VacFVrbFBlVTlKZURCSU9FMVpNblYxUm14M2RURnVlbU5yVGt0Wk1GTkJhVXAzVmt3MVJuQm5NQV8`

// 	// IDトークンをヘッダー・ペイロード・署名に分割
// 	dataArray := strings.Split(idToken, ".")
// 	header, payload, sig := dataArray[0], dataArray[1], dataArray[2]

// 	// header を base64 decode する
// 	headerData, err := base64.RawURLEncoding.DecodeString(header)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return
// 	}

// 	// payload を base64 decode する
// 	payloadData, err := base64.RawURLEncoding.DecodeString(payload)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return
// 	}

// 	// 公開鍵構造体を作る
// 	E := "AQAB"
// 	N := "onV5tzUbqyPfkM6MwUqCtrqun9x20hEUbIUlmAYYuPuMhsaNHJqs1AVzRt2TzaNjmPVddEbU7VMDmeFWUt7vgDi7Xu0leevuIN4VSPbAMGBa0oj9Qopqkn9ePO_7DvIN13ktHgfQqatNBu6uXH6zkUl3VtXnubXrUhx7uyF22dARDc1-pJoj2NnsvgxDRElPMyDkU-siVv3c6cgIEwLEZZPWOcwplPTUB4qeTK6prrPBGQshuE1PWK2ZrYpIvXfzHyEbkGdPnrhcxgCzbKBUFvr8n_sfSurLRoDBLjkURKmgB8T8iRzLyXsCu9D3Hw61LKuex1aeSQLdwOFLuUEBdw"

// 	dn, _ := base64.RawURLEncoding.DecodeString(N)
// 	de, _ := base64.RawURLEncoding.DecodeString(E)

// 	pk := &rsa.PublicKey{
// 		N: new(big.Int).SetBytes(dn),
// 		E: int(new(big.Int).SetBytes(de).Int64()),
// 	}

// 	// 検証するデータ
// 	// ヘッダー+ペイロード部分をハッシュ関数SHA256を使って処理し、メッセージダイジェストを生成
// 	message := sha256.Sum256([]byte(header + "." + payload))

// 	// 署名をbase64 decode する
// 	sigData, err := base64.RawURLEncoding.DecodeString(sig)
// 	if err != nil {
// 		fmt.Println("sig error: ", err)
// 		return
// 	}

// 	if err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, message[:], sigData); err != nil {
// 		fmt.Println("invalid token")
// 	} else {
// 		fmt.Println("valid token")

// 		// JSONを整形して出力
// 		var headerJSON, payloadJSON map[string]interface{}
// 		if err := json.Unmarshal(headerData, &headerJSON); err != nil {
// 			fmt.Println("error", err)
// 			return
// 		}
// 		if err := json.Unmarshal(payloadData, &payloadJSON); err != nil {
// 			fmt.Println("error", err)
// 			return
// 		}

// 		headerIndented, err := json.MarshalIndent(headerJSON, "", "  ")
// 		if err != nil {
// 			fmt.Println("error", err)
// 			return
// 		}
// 		payloadIndented, err := json.MarshalIndent(payloadJSON, "", "  ")
// 		if err != nil {
// 			fmt.Println("error", err)
// 			return
// 		}

// 		fmt.Println("header: ", string(headerIndented))
// 		fmt.Println("payload: ", string(payloadIndented))
// 	}
// }
