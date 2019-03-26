package codes

import "fmt"

var (
	UnFormatParamErr  = &Error{Code: UnFormatParam, Msg: UnFormatParam}
	RecordNotFoundErr = &Error{Code: RecordNotFound, Msg: RecordNotFound}
)

type Error struct {
	Code string `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %s, msg: %s", e.Code, e.Msg)
}
