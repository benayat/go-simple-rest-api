package controller

import (
	"awesomeProject2/db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the homepage")
	fmt.Println("Endpoint hit: homepage")
}


func ReturnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: return all articles")
	json.NewEncoder(w).Encode(db.Articles)
}
 func ReturnSingleArticle(w http.ResponseWriter, r *http.Request){
	 vars:= mux.Vars(r)
	 key:= vars["id"]

	 fmt.Fprintf(w, "Key :"+key)
 }