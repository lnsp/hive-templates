package service

import (
	"github.com/lnsp/hive/lib/service"
	"{{ .Path }}/methods"
)

var Service service.Service

func init() {
	Service = service.New("{{ .Name }}", "{{ .Version }}")
	{{ range .Methods }}
	Service.Register(service.NewMethod("{{ .Name }}", methods.{{ .Name }}Request{}, methods.{{ .Name }}Response{}, methods.{{ .Name }}Handler))
	{{ end }}
}
