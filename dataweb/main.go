package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type products struct {
	id      int
	model   string
	company string
	price   int
}

func main() {
	connStr := "user=postgres password=123456789io dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from productname")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	product := []products{}

	for rows.Next() {
		p := products{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		product = append(product, p)
	}

	for _, p := range product {
		fmt.Println(p.id, p.model, p.company, p.price)
	}
}
