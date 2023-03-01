package main

import (
	"database/sql"
	"fmt"
	"time"
)

func GetPromo(currentDate string) promotion {

	var currentPromo promotion
	currentDateTemp, err := time.Parse("2006-01-02", currentDate)
	checkError(err)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)
	defer db.Close()

	rows, err := db.Query("SELECT promotion_name,start_date,end_date FROM Promotion;")
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		var promoName, startDate, endDate string
		err := rows.Scan(&promoName, &startDate, &endDate)
		checkError(err)

		promoStartDateTemp, err := time.Parse("2006-01-02", startDate[:10])
		checkError(err)
		promoEndDateTemp, err := time.Parse("2006-01-02", endDate[:10])
		checkError(err)
		fmt.Println("debug")

		if (currentDateTemp.After(promoStartDateTemp) && currentDateTemp.Before(promoEndDateTemp)) || (currentDateTemp.Equal(promoStartDateTemp)) || (currentDateTemp.Equal(promoEndDateTemp)) {
			currentPromo.PromoName = promoName
		}
	}

	rows, err = db.Query("SELECT promotion_name,interest_rate FROM Rate;")
	checkError(err)

	for rows.Next() {

		var promoTemp promotion
		err := rows.Scan(&promoTemp.PromoName, &promoTemp.InterestRate)
		checkError(err)
		if currentPromo.PromoName == promoTemp.PromoName {
			currentPromo.InterestRate = promoTemp.InterestRate
			return currentPromo
		}

	}

	return promotion{}
}

func insertAccountDetail(db *sql.DB, Respond respondMessage) error {
	insertInto := fmt.Sprintf(`INSERT INTO Account VALUES('%d','%.2f')`, Respond.ResBody.AccountNumber, Respond.ResBody.InstallmentAmount)
	fmt.Println(insertInto)
	tempResult, err:= db.Exec(insertInto)
	fmt.Println(tempResult)
	checkError(err)
	return nil

}
