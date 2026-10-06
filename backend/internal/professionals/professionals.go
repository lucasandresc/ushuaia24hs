package professionals

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Professional struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Address  string  `json:"address"`
	Phone    string  `json:"phone"`
	Phone2   *string `json:"phone2"`
}

func GetAll(conn *pgx.Conn) ([]Professional, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, name, category, address, phone, phone2 FROM professionals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	professionals := []Professional{}

	for rows.Next() {
		var p Professional
		err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Address, &p.Phone, &p.Phone2)
		if err != nil {
			return nil, err
		}
		professionals = append(professionals, p)
	}
	return professionals, nil
}

func GetByCategory(conn *pgx.Conn, category string) ([]Professional, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, name, category, address, phone, phone2 FROM professionals WHERE category = $1", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	professionals := []Professional{}

	for rows.Next() {
		var p Professional
		err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Address, &p.Phone, &p.Phone2)
		if err != nil {
			return nil, err
		}
		professionals = append(professionals, p)
	}
	return professionals, nil
}

func Create(conn *pgx.Conn, p Professional) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO professionals(name, category, address, phone, phone2) VALUES ($1, $2, $3, $4, $5)", p.Name, p.Category, p.Address, p.Phone, p.Phone2)
	return err
}

func Update(conn *pgx.Conn, id int, p Professional) (int64, error) {
	result, err := conn.Exec(context.Background(), "UPDATE professionals SET name = $1, category = $2, address = $3, phone = $4, phone2 = $5 WHERE id = $6", p.Name, p.Category, p.Address, p.Phone, p.Phone2, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

func Delete(conn *pgx.Conn, id int) (int64, error) {
	result, err := conn.Exec(context.Background(), "DELETE FROM professionals WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
