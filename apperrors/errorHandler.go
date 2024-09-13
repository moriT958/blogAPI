package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/moriT958/go-api/common"
)

// エラーが発生したときのレスポンス処理をここで一括で行う
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError

	// 知らないエラーが出た時の処理
	if !errors.As(err, &appErr) {
		// As(a, b)はaのエラーインターフェースがbに変換可能ならtrueを返す
		// Asを使うのはアサーションをこなうよりいいことがある。
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := common.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	// エラーコードに応じて返すエラーを分岐
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam, PayloadDecodeFailed:
		statusCode = http.StatusBadRequest
	case RequiredAuthorizationHeader, Unauthorizated:
		statusCode = http.StatusUnauthorized
	case NotMatchUser:
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	//  ヘッダに書き込み
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
