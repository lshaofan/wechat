/*
 * 版权所有 (c) 2022 伊犁绿鸟网络科技团队。
 *  response.go  response.go 2022-11-30
 */

package response

type ErrorModel struct {
	Code       int         `json:"errcode"`
	Message    string      `json:"errmsg"`
	Error      interface{} `json:"result"`
	HttpStatus int         `json:"httpStatus" swaggerignore:"true"`
}
type Response struct {
	Code    int         `json:"code" example:"0"`
	Result  interface{} `json:"result"`
	Message string      `json:"message" example:"操作成功"`
}
