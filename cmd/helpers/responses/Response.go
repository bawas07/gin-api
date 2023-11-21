package responses

import (
	// configuration "gin-api/cmd/config"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    *ResponseCodeEnum
	Message *string
	Data    interface{}
}

type ResposeBody struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseReturn struct {
	Data     ResposeBody
	HTTPCode int
}

func ResponseOk(c *gin.Context, data ResponseData) {
	responseCode := ResOK
	data.Code = &responseCode

	if data.Message == nil {
		message := "success"
		data.Message = &message
	}
	ResponseTemplate(c, data)
	return
}

func ResponseTemplate(c *gin.Context, data ResponseData) {
	responseCode := ResOK
	var body ResposeBody
	body.Data = data.Data
	body.Code = responseCode.String().Code
	body.Message = "Success"
	status := responseCode.String().HTTPCode

	if data.Code != nil {
		code := data.Code.String().Code
		body.Code = code
		status = data.Code.String().HTTPCode
	}

	if data.Message != nil {
		body.Message = *data.Message
	}

	c.JSON(status, body)
	return
	// return ResponseReturn{
	// 	Data:     body,
	// 	HTTPCode: status,
	// }
}
