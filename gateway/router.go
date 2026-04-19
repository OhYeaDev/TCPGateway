package gateway

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
)

var routeMap map[string][]string
var routeMu sync.Mutex
var counters = make(map[string]int)

func Load() {
	file, err := os.ReadFile("routes.json")
	if err != nil {
		routeMap = map[string][]string{
			"/api": {"https://httpbin.org"},
		}
		return
	}
	json.Unmarshal(file, &routeMap)
}

func getBackend(path string) string {
	routeMu.Lock()
	defer routeMu.Unlock()

	for prefix, backends := range routeMap {
		if strings.HasPrefix(path, prefix) {
			i := counters[prefix] % len(backends)
			counters[prefix]++
			return backends[i]
		}
	}

	return "https://httpbin.org"
}
