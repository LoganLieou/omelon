package utils

import (
	"database/sql"
	"fmt"
	"log"
)

func UpdateQuantity(db *sql.DB, item_name string, city string, new_quan int) {
	var query string = fmt.Sprintf("UPDATE %s SET quantity=%d WHERE name=%s", city, new_quan, item_name)

	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePrice(db *sql.DB, item_name string, new_price float32) {
	var query string = fmt.Sprintf("UPDATE Inventory SET price=%f", new_price)

	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func PrintWarehouse(db *sql.DB, city string) {
	var query string = "SELECT * FROM " + city

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var item City
		rows.Scan(&item.name, &item.quan, &item.city)
		item.print()
	}
}

