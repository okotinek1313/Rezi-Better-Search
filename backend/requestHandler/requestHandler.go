package requestHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"rezi-better-search/api"
	"rezi-better-search/config"
)

var Port_config config.Config

func RequestHandler() {
	http.HandleFunc("/search", searchHandler)

	existing := config.Read()

	var listener net.Listener
	var port int

	if existing.Port != 0 {
		port = existing.Port
		var err error
		listener, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			listener, port = getListener()
		}
	} else {
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

func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	query := r.URL.Query().Get("q")
	result, err := api.Search(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
