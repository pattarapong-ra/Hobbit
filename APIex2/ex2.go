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

func insertAccountDetail(db *sql.DB, Respond respondMessage) error {
	insertInto := fmt.Sprintf(`INSERT INTO Account VALUES('%d','%.2f')`, Respond.ResBody.AccountNumber, Respond.ResBody.InstallmentAmount)
	fmt.Println(insertInto)
	tempResult, errinsert := db.Exec(insertInto)
	fmt.Println(tempResult)
	if errinsert != nil {
		return errinsert
	}
	return nil

}
