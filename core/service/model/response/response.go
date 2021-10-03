package response

import (
	"net/http"
	"octopus/core/global"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: code, Data: data, Msg: msg})
}

func OK(c *gin.Context) {
	Result(global.SUCCESS_CODE, nil, global.SUCCESS_MESSAGE, c)
}

func OKMessage(msg string, c *gin.Context) {
	Result(global.SUCCESS_CODE, nil, msg, c)
}

func OKData(data interface{}, c *gin.Context) {
	Result(global.SUCCESS_CODE, data, global.SUCCESS_MESSAGE, c)
}

func OKResult(data interface{}, msg string, c *gin.Context) {
	Result(global.SUCCESS_CODE, data, msg, c)
}

func Fail(code int, c *gin.Context) {
	Result(code, nil, global.FAIL_MESSAGE, c)
}

func FailMessage(code int, msg string, c *gin.Context) {
	Result(code, nil, msg, c)
}

func FailDetail(code int, data interface{}, msg string, c *gin.Context) {
	Result(code, data, msg, c)
}
