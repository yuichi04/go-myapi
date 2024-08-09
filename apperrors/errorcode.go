package apperrors

type ErrCode string

const (
	Unknown          ErrCode = "U000" //
	InsertDataFailed ErrCode = "S001" // データベースへのinsert処理に失敗
	GetDataFailed    ErrCode = "S002" // select文の実行に失敗
	NAData           ErrCode = "S003" // 指定されたデータが存在しない
)
