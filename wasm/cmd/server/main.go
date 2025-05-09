package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", "localhost:8080", "listen address to serve")

func main() {
	flag.Parse()
	root, err := os.OpenRoot("site")
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(*addr, http.FileServerFS(root.FS())); err != nil {
		log.Fatal(err)
	}
}
