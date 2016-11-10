package main

import (
	"github.com/xlk3099/urlshortener/web"
)

func main() {
	// Create a new MongoDB session
	// dbSession := web.NewSession("ShortenedURLRecords")
	dbSession := web.NewSession("test")
	// Create a new server using that MongoDB session
	server := web.NewServer(dbSession)
	// Begin listening for HTTP requests
	server.Run()
}
