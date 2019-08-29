package codes

import (
	ferrors "github.com/xgxw/foundation-go/errors"
)

// 错误常量
const (
	Succeed        = "succeed"
	UnFormatParam  = "unformat_param"
	RecordNotFound = "record_not_found"
)

// 错误类别
var (
	UnFormatParamErr  = &ferrors.Error{Code: UnFormatParam, Msg: UnFormatParam}
	RecordNotFoundErr = &ferrors.Error{Code: RecordNotFound, Msg: RecordNotFound}
)
