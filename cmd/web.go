package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Serve static files
	str, _ := os.Getwd()
	fmt.Println(str)
	fs := http.FileServer(http.Dir("./web"))

	// Redirect / to /main.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./web/main.html")
			return
		}
		fs.ServeHTTP(w, r)
	})

	log.Println("Serving on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
