package ex2

import "database/sql"

func PrepareTable2(db *sql.DB) error {
	errCreatetableEmployees := createtableEmployees(db)
	if errCreatetableEmployees != nil {
		return errCreatetableEmployees
	}
	errCreatetableSalaries := createtableSalaries(db)
	if errCreatetableSalaries != nil {
		return errCreatetableSalaries
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
	// +-------------+---------+
	// employee_id is the primary key for this table.
	// Each row of this table indicates the name of the employee whose ID is employee_id.

	dropBeforeCreateTable := `DROP TABLE IF EXISTS Employees ;`
	_, errdropBeforeCreateTable := db.Exec(dropBeforeCreateTable)
	if errdropBeforeCreateTable != nil {
		return errdropBeforeCreateTable
	}
	createTable := `CREATE TABLE Employees (
		xxxxxxxxxxxx xxxx xxxxxx
		);`
	_, errcreateTable := db.Exec(createTable)
	if errcreateTable != nil {
		return errcreateTable
	}
	// ทำการ insert data ลง table Employees ตามข้อมูลด้านล่าง
	// | employee_id | name     |
	// | ----------- | -------- |
	// | 2           | Crew     |
	// | 4           | Haven    |
	// | 5           | Kristian |

	insert := `INSERT INTO xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
	_, errinsert := db.Exec(insert)
	if errinsert != nil {
		return errinsert
	}
	return nil

}

func createtableSalaries(db *sql.DB) error {
	// Table: Salaries
	// +-------------+---------+
	// | Column Name | Type    |
	// +-------------+---------+
	// | employee_id | int     |
	// | salary      | int     |
	// +-------------+---------+
	// employee_id is the primary key for this table.
	// Each row of this table indicates the salary of the employee whose ID is employee_id.

	dropBeforeCreateTable := `DROP TABLE IF EXISTS Salaries ;`
	_, errdropBeforeCreateTable := db.Exec(dropBeforeCreateTable)
	if errdropBeforeCreateTable != nil {
		return errdropBeforeCreateTable
	}

	createTable := `CREATE TABLE Salaries (
		xxxxxxxxxxxx xxxx xxxxxx
		);`
	_, errcreateTable := db.Exec(createTable)
	if errcreateTable != nil {
		return errcreateTable
	}

	// ทำการ insert data ลง table Salaries ตามข้อมูลด้านล่าง
	// | employee_id | salary |
	// | ----------- | ------ |
	// | 5           | 76071  |
	// | 1           | 22517  |
	// | 4           | 63539  |
	insert := `INSERT INTO xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
	_, errinsert := db.Exec(insert)
	if errinsert != nil {
		return errinsert
	}
	return nil

}
