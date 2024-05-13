package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, x); err != nil {
		fmt.Println("Error reading body:", err)
	}
}
