package handlers

import (
	"encoding/json"
	"github.com/spayder/todo-app/database"
	"net/http"
)

func SearchTodos(db database.TodoInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var filter interface{}
		query := request.URL.Query().Get("q")
		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				WriteResponse(writer, http.StatusBadRequest, err)
				return
			}
		}

		res, err := db.Search(filter)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err)
			return
		}

		WriteResponse(writer, http.StatusOK, res)
	}
}
