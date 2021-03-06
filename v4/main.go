package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	dao "./DAO"
	models "./Models"
	"gopkg.in/mgo.v2/bson"
)

var articleDAO = dao.ArticleDAO{}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&article)
	if handleError(err, w) {
		return
	}

	article.Id = bson.NewObjectId()
	err = articleDAO.Insert(article)
	if !handleError(err, w) {
		json.NewEncoder(w).Encode(article)
	}
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := articleDAO.ListAll()
	if !handleError(err, w) {
		json.NewEncoder(w).Encode(articles)
	}
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	article, err := articleDAO.Get(params["id"])

	if !handleError(err, w) {
		json.NewEncoder(w).Encode(article)
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := articleDAO.Remove(params["id"])

	handleError(err, w)
}

func init() {
	articleDAO.Connect()
}

func handleError(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}

	w.WriteHeader(http.StatusOK)
	return false
}

func main() {
	handlerRequest()
}
