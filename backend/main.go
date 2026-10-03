package main

import (
	"context"
	"fmt"

	"github.com/lucasandresc/ushuaia24hs/backend/internal/database"
	"github.com/lucasandresc/ushuaia24hs/backend/internal/professionals"
)

func main() {
	conn, err := database.Connect()
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
	}
	defer conn.Close(context.Background())

	professionalsList, err := professionals.GetAll(conn)
	if err != nil {
		fmt.Println("Error retrieving professionals info:", err)
		return
	}
	for _, value := range professionalsList {
		fullContact := fmt.Sprintf("ID: %d - Name: %s - Category: %s - Address: %s - Phone: %s", value.ID, value.Name, value.Category, value.Address, value.Phone)
		if value.Phone2.Valid {
			fullContact = fullContact + " - " + fmt.Sprintf("Phone 2: %s", value.Phone2.String)
		}
		fmt.Println(fullContact)
	}

}
