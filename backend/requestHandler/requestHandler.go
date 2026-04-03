package requestHandler

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"rezi-better-search/config"
)

var Port_config config.Config

func RequestHandler() {
	existing := config.Read()

	var listener net.Listener
	var port int

	if existing.Port != 0 {
		// use existing port
		port = existing.Port
		var err error
		listener, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			// port is taken get a new one
			listener, port = getListener()
		}
	} else {
		// no port in config get a new one
		listener, port = getListener()
	}

	Port_config = config.Config{Port: port}
	config.AddConfig(Port_config)
	log.Printf("Server running on http://localhost:%d", port)
	log.Fatal(http.Serve(listener, nil))
}

func getListener() (net.Listener, int) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	return listener, listener.Addr().(*net.TCPAddr).Port
}
