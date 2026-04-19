package gateway

import (
	"sync"
	"time"
)

var clients = make(map[string]int)
var mu sync.Mutex

func Allow(ip string) bool {
	mu.Lock()
	defer mu.Unlock()

	clients[ip]++

	go func(ip string) {
		time.Sleep(time.Second)
		mu.Lock()
		clients[ip]--
		mu.Unlock()
	}(ip)

	return clients[ip] <= 10
}
