package router

import (
	"awesomeProject2/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests(){
	MyRouter:=mux.NewRouter().StrictSlash(true)
	MyRouter.HandleFunc("/", controller.HomePage)
	MyRouter.HandleFunc("/all", controller.ReturnAllArticles)
	MyRouter.HandleFunc("/articles/{id}", controller.ReturnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", MyRouter))
}
