package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"virus/utils"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

const (
	Succerr = 0
	Err     = 7
)

func Ok(data any, msg string, c *gin.Context) {
	Result(Succerr, data, msg, c)
}
func OkWithData(data any, c *gin.Context) {
	Result(Succerr, data, "成功", c)
}

func OkWithList(list any, count int64, c *gin.Context) {
	OkWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
}
func OkWithMessage(msg string, c *gin.Context) {
	Result(Succerr, map[string]any{}, msg, c)
}

func OkWith(c *gin.Context) {
	Result(Succerr, map[string]any{}, "成功", c)
}
func Fail(data any, msg string, c *gin.Context) {
	Result(Succerr, data, msg, c)
}

func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Succerr, map[string]any{}, msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Err, map[string]any{}, "未知错误", c)

}