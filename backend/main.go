package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func (s *Server) ProfessionalsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
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
	case http.MethodPost:
		var p professionals.Professional
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(p.Name) == "" || strings.TrimSpace(p.Category) == "" || strings.TrimSpace(p.Address) == "" || strings.TrimSpace(p.Phone) == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}
		err = professionals.Create(s.conn, p)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
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
	http.HandleFunc("/professionals", server.ProfessionalsHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

}
