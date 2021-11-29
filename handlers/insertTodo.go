package handlers

import (
	"encoding/json"
	"github.com/spayder/todo-app/database"
	"github.com/spayder/todo-app/models"
	"io/ioutil"
	"net/http"
)

func InsertTodo(db database.TodoInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		todo := models.Todo{}
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err)
			return
		}

		err = json.Unmarshal(body, &todo)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err)
			return
		}

		res, err := db.Insert(todo)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err)
			return
		}

		WriteResponse(writer, http.StatusOK, res)
	}
}

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}
}
