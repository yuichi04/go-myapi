// ラップ元となるエラーを定義するファイル

package services

import "errors"

// errors.New関数：内部に何もエラーをラップしていない「起点」となるエラーを作ることができる
var ErrNoData = errors.New("get 0 record from db.Query")
