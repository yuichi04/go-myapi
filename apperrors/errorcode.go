package apperrors

type ErrCode string

const (
	Unknown          ErrCode = "U000" // 原因不明
	InsertDataFailed ErrCode = "S001" // データベースへのinsert処理に失敗
	GetDataFailed    ErrCode = "S002" // select文の実行に失敗
	NAData           ErrCode = "S003" // 指定された記事が存在しない
	NoTargetData     ErrCode = "S004" // 指定されたコメント投稿先の記事が存在しない
	UpdateDataFailed ErrCode = "S005" // 指定された記事のいいね数の更新に失敗

	ReqBodyDecodeFailed ErrCode = "R001" // リクエストボディのjsonでコードに失敗
	BadParam            ErrCode = "R002" // リクエストパラメータの値が不正

	RequiredAuthrizationHeader ErrCode = "A001"
	CannotMakeValidator        ErrCode = "A002"
	Unauthrizated              ErrCode = "A003"
	NotMatchUser               ErrCode = "A004"
)
