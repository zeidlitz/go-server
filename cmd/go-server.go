package main

import (
	"fmt"
	"strconv"

	"github.com/zeidlitz/go-server/internal/env"
	"github.com/zeidlitz/go-server/internal/server"
)

func main() {
	hostname := env.GetString("HOSTNAME", "localhost")
	port := env.GetInt("PORT", 8080)
	host := fmt.Sprintf(hostname + ":" + strconv.Itoa(port))
	server.Start(host)
}
