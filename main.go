package main

import (
	"net/http"

	_ "github.com/stealthrocket/net/http"
	"github.com/stealthrocket/net/wasip1"
)

func main() {
	println("server started :3000")

	listener, err := wasip1.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World from WasmEdge!"))
	})

	server := &http.Server{
		Handler: mux,
	}

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
