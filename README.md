TCPGateway

TCPGateway is a lightweight API gateway written in Go. It handles request routing, reverse proxying, rate limiting, and API key validation with minimal overhead.

Features

- Reverse proxy routing
- Config-based routes ("routes.json")
- Round-robin load balancing
- API key authentication
- Basic rate limiting
- Request logging
- Simple web panel

---

Getting Started

Requirements

- Go 1.22+

Setup

git clone https://github.com/OhYeaDev/TCPGateway.git
cd TCPGateway
go mod tidy
go run main.go

Server starts on:

http://localhost:8080

---

Usage

API Request

curl -H "X-API-Key: test123" http://localhost:8080/api/get

---

Web Panel

Open in browser:

http://localhost:8080/panel/

---

Configuration

routes.json

{
  "/api": [
    "https://httpbin.org",
    "https://postman-echo.com"
  ]
}

- Keys = route prefixes
- Values = backend servers
- Requests are distributed using round-robin

---

apikeys.json

[
  "test123",
  "mysecretkey"
]

Requests must include:

X-API-Key: <key>

---

Project Structure

TCPGateway/
├── main.go
├── go.mod
├── routes.json
├── apikeys.json
├── gateway/
│   ├── proxy.go
│   ├── router.go
│   ├── ratelimit.go
│   ├── logger.go
│   └── auth.go
└── web/
    └── index.html

---

Notes

- Default route falls back to "httpbin.org"
- Rate limiting is IP-based
- API keys are loaded at startup
- Designed for extension (auth, metrics, config reload)

---

License

MIT
