package apperrors

type MyAppError struct {
	// ErrCode -> レスポンスとログに表示するエラーコード
	// Message -> レスポンスに表示するエラーメッセージ
	// error -> ログに表示する生の内部エラー

	// ErrCode型のErrCodeフィールド
	// (フィールド名を省略した場合、型名がそのままフィールド名になる)
	ErrCode
	Message string
	Err     error // エラーの入れ子関係のことをエラーチェーンという
}

// errorインターフェース型を満たすために、Errorメソッドを実装
func (myErr *MyAppError) Error() string {
	// return myErr.Message
	return myErr.Err.Error() // 開発中はわかりやすいように内部エラーを直接返すようにする。
}

// errors.Is/errors.Asを使えるようにUnwrapメソッドを定義
// エラーチェーンの中身を取り出すための関数
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
