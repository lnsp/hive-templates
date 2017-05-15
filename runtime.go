package main

import (
	"{{ .Path }}/service"
)

func main() {
	service.Service.Run()
}
