package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spayder/todo-app/config"
	"github.com/spayder/todo-app/database"
	"net/http"
)

func main() {
	conf := config.GetConfig()
	db := database.ConnectDB(conf.Mongo)
	fmt.Println(db)

	router := mux.NewRouter()
	http.ListenAndServe(":6060", router)
}
