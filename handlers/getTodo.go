package handlers

import (
	"github.com/gorilla/mux"
	"github.com/spayder/todo-app/database"
	"net/http"
)

func GetTodo(db database.TodoInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		res, err := db.Get(id)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err)
			return
		}

		WriteResponse(writer, http.StatusOK, res)
	}
}
