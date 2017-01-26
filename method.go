package {{ .Service }}

type {{ .Name }}Request struct {
}

type {{ .Name }}Response struct {
}

var {{ .Name }} service.Method

func {{ .Name }}Handler(request interface{}) (interface{}, error) {
	return &{{ .Name }}Response{}, nil
}

func init() {
	{{ .Name }} = service.NewMethod("{{ .Name }}", {{ .Name }}Request{}, {{ .Name }}Response{}, {{ .Name }}Handler)
	Service.Register({{ .Name }})
}
