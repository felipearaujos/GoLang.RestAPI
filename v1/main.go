package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var articles Articles

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Please insert the correct article Id", 400)
		return
	}

	for i := range articles {
		if articles[i].Id == articleId {
			json.NewEncoder(w).Encode(articles[i])
			return
		}
	}
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	if articles != nil {
		json.NewEncoder(w).Encode(articles)
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Please insert the correct article Id", 400)
		return
	}

	var tempArticles Articles

	for i := range articles {
		if articles[i].Id != articleId {
			tempArticles = append(tempArticles, articles[i])
		}
	}

	articles = tempArticles
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	article.Id = len(articles) + 1
	articles = append(articles, article)

	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Please insert the correct article Id", 400)
		return
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var article Article
	err = json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	for i := range articles {
		if articles[i].Id == articleId {
			articles[i].Title = article.Title
			articles[i].Desc = article.Content
			articles[i].Content = article.Content

			return
		}
	}
}

func main() {
	handlerRequest()
}
