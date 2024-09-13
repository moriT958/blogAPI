package common

import (
	"context"
	"net/http"
)

// コンテキストを扱う処理をcommonとして切り出す
// パッケージ同士の循環参照を防ぐことができる。(apperrorsでmiddlewaresを参照し、middlewaresでapperrorsを参照)
// ハンドラの動作がミドルウェアからのコンテキストに依存しているというアンチパターンも解消できる。

type traceIDKey struct{} // withValueに与えるキー用の独自型

// コンテキストにトレースIDを付加する関数
func SetTraceID(ctx context.Context, traceID int) context.Context {
	// withValueはコンテキストに対して任意のkey-value情報を付与できる。
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// コンテキスト内のトレースIDを取得する関数
func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey{})
	if idInt, ok := id.(int); ok { // Value()の戻り値はany型なのでintにアサーションする必要がある。
		return idInt
	}

	return 0
}

type userNameKey struct{} // コンテキスト中でnameフィールドに対応させるキー構造体

// コンテキストからnameフィールドの値を取り出す関数
func GetUserName(ctx context.Context) string {
	id := ctx.Value(userNameKey{}) // コンテキストからnameフィールドの値を取得
	if usernameStr, ok := id.(string); ok {
		return usernameStr
	}
	return ""
}

// コンテキストにnameフィールドの値をセットする関数
func SetUserName(req *http.Request, name string) *http.Request {
	ctx := req.Context()                              // リクエスト内に含まれるコンテキストを取得
	ctx = context.WithValue(ctx, userNameKey{}, name) // コンテキストにnameフィールドのkey-valueを追加
	req = req.WithContext(ctx)                        // リクエストにコンテキストを追加
	return req
}
