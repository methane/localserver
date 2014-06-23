package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var listen, dir string
	flag.StringVar(&dir, "dir", "", "Directory path for serving.")
	flag.StringVar(&listen, "listen", ":8080", "host:port (default ':8080'")
	flag.Parse()

	if dir == "" {
		dir = filepath.Dir(os.Args[0])
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(listen, http.FileServer(http.Dir(dir))))
}
