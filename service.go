package {{ service.Name }}

import (
	"github.com/lnsp/hive/lib/service"
)

var Service service.Service

func init() {
	Service = service.New("{{ service.Name }}", "0.1.0")
}
