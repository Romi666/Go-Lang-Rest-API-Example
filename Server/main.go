package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type Articles []Article

var article = Articles{
	Article{Title: "ini judul", Desc: "ini deskripsi"},
	Article{Title: "ini judul2", Desc: "ini 2"},
}

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/articles", getArticle)
	http.HandleFunc("/post-article", withLogging(postArticle))
	http.ListenAndServe(":3000", nil)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, "Can't Read Body", http.StatusInternalServerError)
		// }

		// w.Write([]byte(string(body)))

		var newArticle Article
		err := json.NewDecoder(r.Body).Decode(&newArticle)

		if err != nil {
			fmt.Printf("Ada Error bosq")
		}
		var articles = append(article, newArticle)
		json.NewEncoder(w).Encode(&articles)
	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Haloo Golang disini"))
}

func getArticle(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(article)
}

//MiddleWare
func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged koneksi dari", r.RemoteAddr)

		next.ServeHTTP(w, r)
	}
}
