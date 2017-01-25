package {{ service.Name }}

type {{ method.Name }}Request struct {
}

type {{ method.Name }}Response struct {
}

var {{ method.Name }} service.Method

func {{ method.Name }}Handler(request interface{}) (interface{}, error) {
	return &{{ method.Name }}Response{}, nil
}

func init() {
	{{ method.Name }} = service.NewMethod("{{ method.Name }}", {{ method.Name }}Request{}, {{ method.Name }}Response{}, {{ method.Name }}Handler)
	Service.Register({{ method.Name }})
}
