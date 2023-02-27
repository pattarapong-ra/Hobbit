package main

import (
	"database/sql"
	"fmt"
)

func PrepareTable(db *sql.DB) error {
	errCreatetablePromotion := createtablePromotion(db)
	if errCreatetablePromotion != nil {
		return errCreatetablePromotion
	}

	errCreatetableRate := createtableRate(db)
	if errCreatetableRate != nil {
		return errCreatetableRate
	}
	errCreatetableAccount := createtableAccount(db)
	if errCreatetableAccount != nil {
		return errCreatetableAccount
	}
	fmt.Println("done creating table")
	return nil
}
func createtablePromotion(db *sql.DB) error {

	dropBeforeCreateTable := `DROP TABLE IF EXISTS Promotion;`
	_, errdropBeforeCreateTable := db.Exec(dropBeforeCreateTable)
	if errdropBeforeCreateTable != nil {
		return errdropBeforeCreateTable
	}
	createTable := `CREATE TABLE Promotion (
		promotion_name varchar(30),
		description varchar(50),
		start_date date,
		end_date date,
		PRIMARY KEY (promotion_name)
		);`
	_, errcreateTable := db.Exec(createTable)
	if errcreateTable != nil {
		return errcreateTable
	}

	insert := `INSERT INTO Promotion(promotion_name,description,start_date,end_date)
	VALUES('Promo1','Rate < 10','2020-01-01','2020-03-31')`
	_, errinsert := db.Exec(insert)
	insert = `INSERT INTO Promotion(promotion_name,description,start_date,end_date)
	VALUES('Promo2','Rate > 10 < 20','2020-04-01','2020-06-30')`
	_, errinsert = db.Exec(insert)
	insert = `INSERT INTO Promotion(promotion_name,description,start_date,end_date)
	VALUES('Promo3','Rate > 20 < 28','2020-07-01','2020-12-30')`
	_, errinsert = db.Exec(insert)

	if errinsert != nil {
		return errinsert
	}
	return nil

}

func createtableRate(db *sql.DB) error {

	dropBeforeCreateTable := `DROP TABLE IF EXISTS Rate;`
	_, errdropBeforeCreateTable := db.Exec(dropBeforeCreateTable)
	if errdropBeforeCreateTable != nil {
		return errdropBeforeCreateTable
	}

	createTable := `CREATE TABLE Rate (
		rate varchar(20),
		interest_rate float,
		promotion_name varchar(20),
		PRIMARY KEY (rate)
		);`
	_, errcreateTable := db.Exec(createTable)
	if errcreateTable != nil {
		return errcreateTable
	}

	insert := `INSERT INTO Rate(rate,interest_rate,promotion_name)
	VALUES('RatePromo1','2.5','Promo1')`
	_, errinsert := db.Exec(insert)
	insert = `INSERT INTO Rate(rate,interest_rate,promotion_name)
	VALUES('RatePromo2','18','Promo2')`
	_, errinsert = db.Exec(insert)
	insert = `INSERT INTO Rate(rate,interest_rate,promotion_name)
	VALUES('RatePromo3','25','Promo3')`
	_, errinsert = db.Exec(insert)

	if errinsert != nil {
		return errinsert
	}
	return nil

}
func createtableAccount(db *sql.DB) error {

	dropBeforeCreateTable := `DROP TABLE IF EXISTS Account;`
	_, errdropBeforeCreateTable := db.Exec(dropBeforeCreateTable)
	if errdropBeforeCreateTable != nil {
		return errdropBeforeCreateTable
	}

	createTable := `CREATE TABLE Account (
		account_number int,
		installment_amount float,
		PRIMARY KEY (account_number)
		);`
	_, errcreateTable := db.Exec(createTable)
	if errcreateTable != nil {
		return errcreateTable
	}

	/*insert := `INSERT INTO Salaries(employee_id,salary)
	VALUES('5','76071')`
	_, errinsert := db.Exec(insert)
	insert = `INSERT INTO Salaries(employee_id,salary)
	VALUES('1','22517')`
	_, errinsert = db.Exec(insert)
	insert = `INSERT INTO Salaries(employee_id,salary)
	VALUES('4','63539')`
	_, errinsert = db.Exec(insert)

	if errinsert != nil {
		return errinsert
	}*/
	return nil

}
