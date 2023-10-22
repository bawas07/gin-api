package configuration

import "net/http"

type ResponseCodeEnum int64

const (
	ResOK ResponseCodeEnum = iota
	ResServerError
	ResNoRoute
	ResPanicError
)

type ResponseCode struct {
	Code     string
	HTTPCode int
}

func (code ResponseCodeEnum) String() ResponseCode {
	switch code {
	case ResOK:
		return ResponseCode{
			Code:     "OK",
			HTTPCode: http.StatusOK,
		}
	case ResServerError:
		return ResponseCode{
			Code:     "E_SERVER",
			HTTPCode: http.StatusInternalServerError,
		}
	case ResNoRoute:
		return ResponseCode{
			Code:     "E_NOROUTE",
			HTTPCode: http.StatusNotFound,
		}
	case ResPanicError:
		return ResponseCode{
			Code:     "E_PANIC",
			HTTPCode: http.StatusInternalServerError,
		}
	}

	return ResponseCode{
		Code:     "E_UNKNOWN",
		HTTPCode: http.StatusInternalServerError,
	}
}
