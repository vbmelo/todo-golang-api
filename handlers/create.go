package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vbmelo/todo-golang-api/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error while JSON decoding: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]interface{} // Use interface{} instead of any
	if err != nil {
		resp = map[string]interface{}{
			"Error":   true,
			"Message": fmt.Sprintf("Error while inserting: %v", err),
		}
	} else {
		resp = map[string]interface{}{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserted with success! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
