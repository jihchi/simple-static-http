package main

import (
	"flag"
	"log"
	"net/http"

	httpgzip "github.com/daaku/go.httpgzip"
	"github.com/yookoala/realpath/realpath"
)

func main() {
	var err error
	var addr string
	var dir string
	var serve string

	flag.StringVar(&addr, "address", ":8080", "Specify serving HTTP address")
	flag.StringVar(&dir, "path", "./", "Specify serving directory")
	flag.StringVar(&serve, "serve", "/", "Specify HTTP serving path")
	flag.Parse()

	dir, err = realpath.Realpath(dir)
	if err != nil {
		log.Fatal("Failed to get serving directory:", err)
	}

	http.Handle(serve, httpgzip.NewHandler(http.FileServer(http.Dir(dir))))
	log.Printf("Listening \"http://localhost%s%s\"", addr, serve)
	log.Printf("Serving directory: \"%s\"", dir)
	log.Println("Press `ctrl-c` keys to close server...")
	log.Fatal("Failed to listening and serving:", http.ListenAndServe(addr, nil))
}
