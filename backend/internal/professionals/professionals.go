package professionals

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Professional struct {
	ID       int
	Name     string
	Category string
	Address  string
	Phone    string
	Phone2   pgtype.Text
}

func GetAll(conn *pgx.Conn) ([]Professional, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, name, category, address, phone, phone2 FROM professionals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var professionals []Professional

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
