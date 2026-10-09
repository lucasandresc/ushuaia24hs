package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lucasandresc/ushuaia24hs/backend/internal/database"
	"github.com/lucasandresc/ushuaia24hs/backend/internal/handlers"
)

func main() {
	conn, err := database.Connect()
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
	}
	defer conn.Close(context.Background())
	server := handlers.NewServer(conn)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /professionals", server.GetProfessionals)
	mux.HandleFunc("POST /professionals", server.CreateProfessional)
	mux.HandleFunc("GET /professionals/{id}", server.GetProfessionalByID)
	mux.HandleFunc("PUT /professionals/{id}", server.UpdateProfessional)
	mux.HandleFunc("DELETE /professionals/{id}", server.DeleteProfessional)

	err = http.ListenAndServe(":8080", handlers.CorsMiddleware(mux))
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
