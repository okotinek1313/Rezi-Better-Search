package main

import (
	"rezi-better-search/config"
	"rezi-better-search/requestHandler"
)

func main() {
	config.Create()
	requestHandler.RequestHandler()
}
