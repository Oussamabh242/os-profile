package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	fs := http.FileServer(http.Dir("./web"))

	// INDEX
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./web/index.html")
		}

		fs.ServeHTTP(w, r)
	})

	// Fetch Post_By_title
	http.HandleFunc("GET /posts/{post_title}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside /posts/{}")
		postTitle := r.PathValue("post_title")
		file, err := os.ReadFile("./static/posts/" + postTitle + ".md")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(file)
	})

	// Get All Posts w/ filters search,

	http.HandleFunc("GET /posts", func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")

		postsDir, err := os.ReadDir("./static/posts/")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var allPosts []string = []string{}
		for _, entr := range postsDir {
			if search == "" {
				allPosts = append(allPosts, entr.Name())
				continue
			}
			if strings.Contains(entr.Name(), search) {
				allPosts = append(allPosts, entr.Name())
			}
		}
		w.Header().Set("Content-Type", "text/json")
		j, err := json.Marshal(map[string]interface{}{
			"posts": allPosts,
		})
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(j)
		return

	})

	log.Println("Serving on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
