package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"systolic/database"
	"systolic/models"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("🔧 Creating a new todo", r.Body)
	var newTodo models.Todo
	log.Println("🔧 Decoding request body into Todo model", newTodo)
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "❌ Invalid request payload", http.StatusBadRequest)
		return
	}

	res, _, err := database.SupaClient.From("todos").Insert(newTodo, false, " ", " ", "").Execute()

	if err != nil {
		log.Println("❌ Error inserting todo:", err)
		http.Error(w, "❌ Insert failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
