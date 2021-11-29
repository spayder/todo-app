package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spayder/todo-app/config"
	"github.com/spayder/todo-app/database"
	"github.com/spayder/todo-app/handlers"
	"net/http"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()
	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	client := &database.TodoClient{
		Col: collection,
		Ctx: ctx,
	}

	fmt.Println(db)

	router := mux.NewRouter()
	router.HandleFunc("/todos", handlers.SearchTodos(client)).Methods("GET")
	router.HandleFunc("/todos/{id}", handlers.GetTodo(client)).Methods("GET")
	router.HandleFunc("/todos", handlers.InsertTodo(client)).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodo(client)).Methods("PATCH")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo(client)).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
