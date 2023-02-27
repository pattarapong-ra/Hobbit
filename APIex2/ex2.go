package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func GetPromo(currentDate string) float64 {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()


	rows, err := db.Query("SELECT employee_id FROM Employees;")
	if err != nil {
		return -1
	}
	defer rows.Close()

	var employees []int
	fmt.Println("Employee ID")
	for rows.Next() {
		var emp int
		if err := rows.Scan(&emp); err != nil {
			fmt.Println(err)
		}
		fmt.Println(emp)
		employees = append(employees, emp)

	}

	return -1
}
