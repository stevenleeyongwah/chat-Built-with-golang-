package main

import (
	"database/sql"
	"log"
	"net/http"
	"go-chat/controllers"
	_ "github.com/lib/pq"
	"go-chat/storage"
)

func main() {
	// Connect to PostgreSQL
	db, err := sql.Open("postgres", "postgres://root:root@db:5432/chatdb?sslmode=disable")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// ✅ Assign to shared global in storage package
	storage.DB = db

	// ✅ Create and start the hub
	hub := NewHub()
	go hub.Run()

	// ✅ Use ServeWs to handle WebSocket connections
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})
	http.HandleFunc("/messages", controllers.GetMessages)
	http.HandleFunc("/groups", controllers.ListGroups)

	// ✅ Start server
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
