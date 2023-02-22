package main

import (
	"database/sql"
	"fmt"
)

func PrepareTable1(db *sql.DB) error {
	err := createtableEmployees(db)
	if err != nil {
		return err
	}
	return nil
}
func createtableEmployees(db *sql.DB) error {
	// 	Table: Employees
	// +-------------+---------+
	// | Column Name | Type    |
	// +-------------+---------+
	// | employee_id | int     |
	// | name        | varchar |
	// | salary      | int     |
	// +-------------+---------+
	// employee_id is the primary key for this table.
	// Each row of this table indicates the employee ID, employee name, and salary.
	fmt.Println("abcdeg")
	dropBeforeCreateTable := `DROP TABLE IF EXISTS Employees ;`
	_, errdropBeforeCreateTable := db.Exec(dropBeforeCreateTable)
	if errdropBeforeCreateTable != nil {
		return errdropBeforeCreateTable
	}
	createTable := `CREATE TABLE Employees (
		employee_id int,
		name varchar(255),
		salary int
		);`
	_, errcreateTable := db.Exec(createTable)
	if errcreateTable != nil {
		return errcreateTable
	}
	// ทำการ insert data ลง table Employees ตามข้อมูลด้านล่าง
	// Employees table:
	// +-------------+---------+--------+
	// | employee_id | name    | salary |
	// +-------------+---------+--------+
	// | 2           | mEir    | 3000   |
	// | 3           | MichAel | 3800   |
	// | 7           | adDilyn | 7400   |
	// | 8           | JuaN    | 6100   |
	// | 9           | kannon  | 7700   |
	// +-------------+---------+--------+

	insert := `INSERT INTO Employees(employee_id,name,salary)
	VALUES('2','mEir','3000')`
	_, errinsert := db.Exec(insert)
	insert = `INSERT INTO Employees(employee_id,name,salary)
	VALUES('3','MichAel','3800')`
	_, errinsert = db.Exec(insert)
	insert = `INSERT INTO Employees(employee_id,name,salary)
	VALUES('7','adDilyn','7400')`
	_, errinsert = db.Exec(insert)
	insert = `INSERT INTO Employees(employee_id,name,salary)
	VALUES('8','JuaN','6100')`
	_, errinsert = db.Exec(insert)
	insert = `INSERT INTO Employees(employee_id,name,salary)
	VALUES('9','kannon','7700')`
	_, errinsert = db.Exec(insert)
	if errinsert != nil {
		return errinsert
	}
	return nil

}
