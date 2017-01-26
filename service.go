package {{ .Name }}

import (
	"github.com/lnsp/hive/lib/service"
)

var Service service.Service

func init() {
	Service = service.New("{{ .Name }}", "{{ .Version }}")
}
