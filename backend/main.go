package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func (s *Server) CreateProfessional(w http.ResponseWriter, r *http.Request) {
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
}

func (s *Server) UpdateProfessional(w http.ResponseWriter, r *http.Request) {
	var p professionals.Professional
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(p.Name) == "" || strings.TrimSpace(p.Category) == "" || strings.TrimSpace(p.Address) == "" || strings.TrimSpace(p.Phone) == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	rows, err := professionals.Update(s.conn, id, p)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if rows == 0 {
		http.Error(w, "Professional not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) DeleteProfessional(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	rows, err := professionals.Delete(s.conn, id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if rows == 0 {
		http.Error(w, "Professional not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	conn, err := database.Connect()
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
	}
	defer conn.Close(context.Background())
	server := &Server{conn: conn}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", Home)
	mux.HandleFunc("GET /professionals", server.GetProfessionals)
	mux.HandleFunc("POST /professionals", server.CreateProfessional)
	mux.HandleFunc("PUT /professionals/{id}", server.UpdateProfessional)
	mux.HandleFunc("DELETE /professionals/{id}", server.DeleteProfessional)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
