/*
 * 版权所有 (c) 2022 伊犁绿鸟网络科技团队。
 *  response.go  response.go 2022-11-30
 */

package response

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// CommonError 微信返回的通用错误json
type CommonError struct {
	apiName string
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (c *CommonError) Error() string {
	return fmt.Sprintf("%s Error , errcode=%d , errmsg=%s", c.apiName, c.ErrCode, c.ErrMsg)
}

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

// DecodeWithError 将返回值按照解析
func DecodeWithError(response []byte, obj interface{}, apiName string) error {
	err := json.Unmarshal(response, obj)
	if err != nil {
		return fmt.Errorf("json Unmarshal Error, err=%v", err)
	}
	responseObj := reflect.ValueOf(obj)
	if !responseObj.IsValid() {
		return fmt.Errorf("obj is invalid")
	}
	commonError := responseObj.Elem().FieldByName("CommonError")
	if !commonError.IsValid() || commonError.Kind() != reflect.Struct {
		return fmt.Errorf("commonError is invalid or not struct")
	}
	errCode := commonError.FieldByName("ErrCode")
	errMsg := commonError.FieldByName("ErrMsg")
	if !errCode.IsValid() || !errMsg.IsValid() {
		return fmt.Errorf("errcode or errmsg is invalid")
	}
	if errCode.Int() != 0 {
		return &CommonError{
			apiName: apiName,
			ErrCode: errCode.Int(),
			ErrMsg:  errMsg.String(),
		}

	}
	return nil
}
