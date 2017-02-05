package main

import (
	"encoding/json"
	"fmt"

	"{{.Path}}"
)

func main() {
	output, err := json.MarshalIndent({{.Name}}.Service, "", "  ")
	if err != nil {
		fmt.Println("Failed to retrieve service information.")
		return
	}
	fmt.Println(string(output))
}
