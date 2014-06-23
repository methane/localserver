package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var listen string
	flag.StringVar(&listen, "listen", ":8080", "host:port (default ':8080'")
	flag.Parse()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(dir))))
}
