package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

// 🔗 Embed the static directory
//
//go:embed static/*
var embeddedFiles embed.FS

func main() {
	// 🎛️ CLI flags
	port := flag.String("port", "8080", "Port to run the web server on")
	flag.Parse()

	// 🗂️ Create file system for /static/*
	staticFS, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	// 📦 Serve embedded files
	http.Handle("/", http.FileServer(http.FS(staticFS)))

	// 🌐 Start server
	addr := fmt.Sprintf(":%s", *port)
	fmt.Printf("🚀 Serving embedded files on http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
