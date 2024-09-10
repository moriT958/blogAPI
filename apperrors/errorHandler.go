package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
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

	var statusCode int

	// エラーコードに応じて返すエラーを分岐
	switch appErr.ErrCode {
	case NAData, BadParam:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	//  ヘッダに書き込み
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
