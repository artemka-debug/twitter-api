package utils

import (
	"encoding/json"
	"fmt"
)

func PrintMap(x map[string]interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}
