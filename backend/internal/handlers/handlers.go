package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/lucasandresc/ushuaia24hs/backend/internal/professionals"
)

type Server struct {
	conn *pgx.Conn
}

func NewServer(conn *pgx.Conn) *Server {
	return &Server{conn: conn}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main Page")
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

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
