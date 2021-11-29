package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/spayder/todo-app/database"
	"io/ioutil"
	"net/http"
)

func UpdateTodo(db database.TodoInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err)
			return
		}

		var todo interface{}
		err = json.Unmarshal(body, &todo)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err)
			return
		}

		res, err := db.Update(id, todo)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err)
			return
		}

		WriteResponse(writer, http.StatusOK, res)
	}
}
