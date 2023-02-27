package main

import (
	"database/sql"
	"fmt"
	"time"
)

func GetPromo(currentDate string) promotion {

	currentDateTemp, errDate := time.Parse("2006-01-02", currentDate)
	if errDate != nil {
		fmt.Println(errDate)
	}

	var currentPromo promotion

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()

	rows, err := db.Query("SELECT promotion_name,start_date,end_date FROM Promotion;")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	/*var promotionList []promotion
	for rows.Next() {
		var promoTemp promotion
		if err := rows.Scan(&promoTemp.PromoName, promoTemp.Start_date, promoTemp.End_date); err != nil {
			fmt.Println(err)
		}
		promotionList = append(promotionList, promoTemp)
	}*/
	for rows.Next() {
		var promoName, startDate, endDate string
		if err := rows.Scan(&promoName, &startDate, &endDate); err != nil {
			fmt.Println(err)
		}

		promoStartDateTemp, errDate := time.Parse("2006-01-02", startDate[:10])
		if errDate != nil {
			fmt.Println(errDate)
		}
		promoEndDateTemp, errDate := time.Parse("2006-01-02", endDate[:10])
		if errDate != nil {
			fmt.Println(errDate)
		}

		if (currentDateTemp.After(promoStartDateTemp) && currentDateTemp.Before(promoEndDateTemp)) || (currentDateTemp.Equal(promoStartDateTemp)) || (currentDateTemp.Equal(promoEndDateTemp)) {
			currentPromo.PromoName = promoName
		}
	}
	/*for _, val := range promotionList {
		if currentDate >= val.Start_date && currentDate <= val.End_date{
			currentPromo = val.PromoName
		}
	}*/

	rows, err = db.Query("SELECT promotion_name,interest_rate FROM Rate;")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {

		var promoTemp promotion
		if err := rows.Scan(&promoTemp.PromoName, &promoTemp.InterestRate); err != nil {
			fmt.Println(err)
		}
		if currentPromo.PromoName == promoTemp.PromoName {
			currentPromo.InterestRate = promoTemp.InterestRate
			return currentPromo
		}

	}

	return promotion{}
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
func insertAccountDetail(db *sql.DB, installmentRespond respondMessage) error {
	insertInto := fmt.Sprintf("INSERT INTO account(account_number,installment_amount) VALUES('%d','%.2f');", installmentRespond.ResBody.AccountNumber, installmentRespond.ResBody.InstallmentAmount)
	//insertInto := fmt.Sprintf("INSERT INTO account(account_number,installment_amount) VALUES('")+installmentRespond.ResBody.AccountNumber+installmentRespond.ResBody.InstallmentAmount)
	fmt.Println(insertInto)
	_, errinsert := db.Exec(insertInto)
	if errinsert != nil {
		return errinsert
	}
	return nil

}
