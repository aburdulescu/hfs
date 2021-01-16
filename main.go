package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var dir string
	var addr string
	flag.StringVar(&dir, "dir", ".", "path to directory which contains files")
	flag.StringVar(&addr, "addr", "localhost:12345", "listen address(address:port)")
	flag.Parse()

	fileServer := http.FileServer(http.Dir(dir))
	http.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(addr, nil))
}
