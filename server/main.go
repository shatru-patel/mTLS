package main

import (
	// "fmt"
	// "io"

	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Println("This is an example server.")
	w.Write([]byte("This is an example server TLS.\n"))
}

func main() {
	certPath := "../certs/server.crt"
	keyPath := "../certs/server.key"
	fmt.Println("Radsec Server has starting")
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":2083", certPath, keyPath, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
