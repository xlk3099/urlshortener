package main

import (
	"github.com/xlk3099/urlshortener/web"
)

func main() {
	// Create a new MongoDB session
	dbSession := web.NewSession("test")
	defer dbSession.Close()

	// Create a new server using that MongoDB session
	server := web.NewServer(dbSession)

	// Begin listening for HTTP requests
	server.Run()
}
