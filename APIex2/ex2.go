package main

import (
	"database/sql"
	"fmt"
)

func GetPromo(currentDate string) float64 {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()

	rows, err := db.Query("SELECT promotion_name,start_date,end_date FROM Promotion;")
	if err != nil {
		return -1
	}
	defer rows.Close()

	var promotionList []promotion
	for rows.Next() {
		var promoTemp promotion
		if err := rows.Scan(&promoTemp.PromoName, promoTemp.Start_date, promoTemp.End_date); err != nil {
			fmt.Println(err)
		}
		promotionList = append(promotionList, promoTemp)
	}

	rows, err = db.Query("SELECT promotion_name,interest_rate FROM Rate;")
	if err != nil {
		return -1
	}
	defer rows.Close()

	var rateList []rate
	for rows.Next() {
		var rateTemp rate
		if err := rows.Scan(&rateTemp.PromoName, rateTemp.InterestRate); err != nil {
			fmt.Println(err)
		}
		rateList = append(rateList, rateTemp)
	}


	return -1
}

/*var employeeBonus []EmployeeQuery
var employees []Employee
for rows.Next() {
	var emp Employee
	if err := rows.Scan(&emp.EmployeeID, &emp.Name, &emp.Salary); err != nil {
		fmt.Println(err)
	}
	employees = append(employees, emp)
}
for _, val := range employees {
	var tempEmp EmployeeQuery
	tempEmp.EmployeeID = val.EmployeeID
	if val.Name[0:1] != "M" && (val.EmployeeID%2) == 1 {
		tempEmp.Bonus = val.Salary
	} else {
		tempEmp.Bonus = 0
	}
	employeeBonus = append(employeeBonus, tempEmp)
}*/
