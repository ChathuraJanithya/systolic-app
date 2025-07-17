package main

import (
	"log"
	"net/http"
	"systolic/config"
	"systolic/database"
	"systolic/routes"
)

func main() {
    config.LoadConfig()
	log.Println("🔧 Loaded configuration")

	database.ConnectDB()
	log.Println("🔧 Database connection established")

	r := routes.SetupRoutes()

	log.Println("🚀 Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

