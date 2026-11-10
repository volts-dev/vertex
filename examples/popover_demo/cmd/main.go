package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Identify root directory (where static and template folders reside)
	root := "."
	if _, err := os.Stat("template"); os.IsNotExist(err) {
		// Try parent directory if running from cmd subfolder
		if _, err := os.Stat("../template"); err == nil {
			root = ".."
		}
	}

	staticDir := filepath.Join(root, "static")
	templateDir := filepath.Join(root, "template")

	log.Printf("Serving static from: %s", staticDir)
	log.Printf("Serving template from: %s", templateDir)

	// Serve static files
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve index.html at root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(templateDir, "index.html"))
	})

	log.Println("Listening on :9998...")
	err := http.ListenAndServe(":9998", nil)
	if err != nil {
		log.Fatal(err)
	}
}
