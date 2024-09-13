package apperrors

type ErrCode string

// エラーコードを定義するための定数
const (
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"

	RequiredAuthorizationHeader ErrCode = "A001"
	CannotMakeValidator         ErrCode = "A002"
	Unauthorizated              ErrCode = "A003"
	NotMatchUser                ErrCode = "A004"
	PayloadDecodeFailed         ErrCode = "A005"
)

// 受け取ったエラーにエラーコードを含めた形でラップする関数
func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
