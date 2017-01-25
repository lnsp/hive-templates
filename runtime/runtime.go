package main

import (
	{{ service.Path }}{{ service.Name }}
)

func main() {
	{{ service.Name }}.Service.Run()
}
