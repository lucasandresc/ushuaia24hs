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
	category := r.URL.Query().Get("category")
	var professionalsList []professionals.Professional
	var err error
	if category == "" {
		professionalsList, err = professionals.GetAll(s.conn)
	} else {
		professionalsList, err = professionals.GetByCategory(s.conn, category)
	}

	if err != nil {
		fmt.Println("Error getting professionals:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
