package methods

import (
	"github.com/lnsp/hive/lib/service"
)

type {{ .Name }}Request struct {
	Message string
}

type {{ .Name }}Response struct {
	Message string
}

var {{ .Name }} service.Method

func {{ .Name }}Handler(requestData interface{}) (interface{}, *service.Error) {
	request := requestData.(*{{ .Name }}Request)
	return &{{ .Name }}Response{ Message: "Received message '" + request.Message + "'" }, nil
}
