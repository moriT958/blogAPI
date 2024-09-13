package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/moriT958/go-api/apperrors"
	"github.com/moriT958/go-api/common"
)

// gcp側で設定したClientID
// const (
// 	googleClientId = "346884273784-h1f2kbn40frosj0ssffggrombo2h25o9.apps.googleusercontent.com"
// )

// 認証のミドルウェア
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// ヘッダからAuthorizationフィールドを抜き出す
		authorization := req.Header.Get("Authorization")

		// Authorizationフィールドが"Bearer [IDトークン]"の形になっているか検証
		authHeaders := strings.Split(authorization, " ")

		// 空白区切りで 2 つに分かれるか
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// 空白区切りで分け,1つ目がBearer,2つ目が空ではないか
		bearerStr, accessToken := authHeaders[0], authHeaders[1]
		if bearerStr != "Bearer" || accessToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		tokenArray := strings.Split(accessToken, ".")
		_, payloadEncoded, _ := tokenArray[0], tokenArray[1], tokenArray[2]
		payloadStr, err := base64.RawURLEncoding.DecodeString(payloadEncoded)
		if err != nil {
			err = apperrors.PayloadDecodeFailed.Wrap(err, "failed to decode payload")
			apperrors.ErrorHandler(w, req, err)
			return
		}
		var payloadMap map[string]any // map型は動的な構造体のようなもの。dictのようなもの。
		if err := json.Unmarshal(payloadStr, &payloadMap); err != nil {
			err = apperrors.PayloadDecodeFailed.Wrap(err, "failed to parse json")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// "name"フィールドを取得して表示
		name, ok := payloadMap["name"] // valueはany型で宣言しているのでstringにアサーションする必要がある
		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token name")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// contextにユーザー名をセット
		req = common.SetUserName(req, name.(string))

		// 本物のハンドラへ
		next.ServeHTTP(w, req)
	})
}
