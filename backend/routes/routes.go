package routes

import (
	"systolic/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    r := mux.NewRouter()
	
	r.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")

    return r
}
