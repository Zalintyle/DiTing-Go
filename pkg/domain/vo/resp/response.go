package resp

import (
	"DiTing-Go/pkg/domain/enum"

	"github.com/gin-gonic/gin"
)

// ResponseData 表示统一响应的JSON格式
type ResponseData struct {
	Success bool        `json:"success"` // 是否成功
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

// ErrorResponse 是一个辅助函数，用于创建错误响应
// 参数：
//
//	c *gin.Context：Gin上下文对象，用于处理HTTP请求和响应
//	code int：500
//	message string：响应消息，用于描述响应的错误信息或提示信息
func ErrorResponse(c *gin.Context, message string) {
	c.JSON(enum.ERROR, ResponseData{
		Code:    enum.ERROR,
		Success: false,
		Message: message,
		Data:    nil,
	})
}

// SuccessResponse 是一个辅助函数，用于创建成功响应
// 参数：
//
//	c *gin.Context：Gin上下文对象，用于处理HTTP请求和响应
//	code int：200
//	data interface{}：响应数据，用于描述请求处理成功后返回的具体数据
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(enum.SUCCESS, ResponseData{
		Code:    enum.SUCCESS,
		Success: true,
		Message: "success",
		Data:    data,
	})
}

func SuccessResponseWithMsg(c *gin.Context, msg string) {
	c.JSON(enum.SUCCESS, ResponseData{
		Code:    enum.SUCCESS,
		Success: true,
		Message: msg,
		Data:    nil,
	})
}
func ReturnSuccessResponse(c *gin.Context, response ResponseData) {
	c.JSON(enum.SUCCESS, response)
}
func ReturnErrorResponse(c *gin.Context, response ResponseData) {
	c.JSON(enum.ERROR, response)
}

// ErrorResponseData 是一个辅助函数，用于创建错误响应
func ErrorResponseData(msg string) ResponseData {
	return ResponseData{
		Code:    enum.ERROR,
		Success: false,
		Message: msg,
		Data:    nil,
	}
}

// SuccessResponseData 是一个辅助函数，用于创建成功响应
func SuccessResponseData(data interface{}) ResponseData {
	return ResponseData{
		Code:    enum.SUCCESS,
		Success: true,
		Message: "success",
		Data:    data,
	}
}

// SuccessResponseDataWithMsg 是一个辅助函数，用于创建成功响应
func SuccessResponseDataWithMsg(msg string) ResponseData {
	return ResponseData{
		Code:    enum.SUCCESS,
		Success: true,
		Message: msg,
		Data:    nil,
	}
}
