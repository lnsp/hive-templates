package {{ .Service }}

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

func {{ .Name }}Handler(requestData interface{}) (interface{}, error) {
	request := requestData.(*{{ .Name }}Request)
	return &{{ .Name }}Response{ Message: "Received message '" + request.Message + "'" }, nil
}

func init() {
	{{ .Name }} = service.NewMethod("{{ .Name }}", {{ .Name }}Request{}, {{ .Name }}Response{}, {{ .Name }}Handler)
	Service.Register({{ .Name }})
}
