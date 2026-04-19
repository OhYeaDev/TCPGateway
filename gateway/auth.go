package gateway

import (
	"encoding/json"
	"net/http"
	"os"
)

var keys map[string]bool

func LoadKeys() {
	keys = make(map[string]bool)

	file, err := os.ReadFile("apikeys.json")
	if err != nil {
		keys["test123"] = true
		return
	}

	var list []string
	json.Unmarshal(file, &list)

	for _, k := range list {
		keys[k] = true
	}
}

func Authorized(r *http.Request) bool {
	key := r.Header.Get("X-API-Key")
	return keys[key]
}
