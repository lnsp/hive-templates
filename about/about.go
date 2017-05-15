package main

import (
	"encoding/json"
	"fmt"

	"{{.Path}}/service"
)

func main() {
	output, err := json.MarshalIndent(service.Service, "", "  ")
	if err != nil {
		fmt.Println("Failed to retrieve service information.")
		return
	}
	fmt.Println(string(output))
}
