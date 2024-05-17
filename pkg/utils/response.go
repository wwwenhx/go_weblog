package utils

import (
	"gin_gomicro/pkg/e"
)

type Response struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func (r *Response) ResponseSuccess(data interface{}) *Response {
	r.Code = e.Success
	r.Data = data
	r.Message = "Success"
	return r
}

func (r *Response) ResponseFail(data interface{}, message string) *Response {
	r.Code = e.Error
	r.Message = message
	r.Data = data
	return r
}
