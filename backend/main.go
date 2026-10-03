package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/lucasandresc/ushuaia24hs/backend/internal/database"
	"github.com/lucasandresc/ushuaia24hs/backend/internal/professionals"
)

type Server struct {
	conn *pgx.Conn
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main page for testing")
}

func (s *Server) GetProfessionals(w http.ResponseWriter, r *http.Request) {
	professionalsList, err := professionals.GetAll(s.conn)
	if err != nil {
		fmt.Println("Error getting professionals:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(professionalsList)
}

func main() {
	conn, err := database.Connect()
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
	}
	defer conn.Close(context.Background())
	server := &Server{conn: conn}

	http.HandleFunc("/", Home)
	http.HandleFunc("/professionals", server.GetProfessionals)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

}
