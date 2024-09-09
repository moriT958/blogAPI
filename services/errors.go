package services

import "errors"

// ラップする起点にできるエラーを定義しておく
var ErrNoData = errors.New("get 0 record from db.Query")
