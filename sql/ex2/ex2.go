package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

// 1. (union) Write an SQL query to report the IDs of all the employees with missing information.
// The information of an employee is missing if:
// 		The employee's name is missing, or
// 		The employee's salary is missing.
// Return the result table ordered by employee_id in ascending order.

// The query result format is in the following example.
// +-------------+
// | employee_id |
// +-------------+
// | 1           |
// | 2           |
// +-------------+
//ตาราง Employees employee_id name
//ตาราง Salaries employee_id salary
//รวมตารางแล้วเอาตัวที่ชื่อโบ๋หรือเงินโบ๋
//SELECT employee_id FROM Employees WHERE name IS NULL UNION SELECT employee_id FROM Salaries WHERE salary IS NULL ORDER BY employee_ID
// getEmployeesWithMissingInformation : Employees With Missing Information
func GetEmployeesWithMissingInformation(w http.ResponseWriter, r *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()
/*SELECT employee_id
FROM Employees
UNION
SELECT employee_id
FROM Salaries
ORDER BY employee_id;*/
	rows, err := db.Query("select Employees.employee_id FROM Employees LEFT JOIN Salaries on Employees.employee_id = Salaries.employee_id where Salaries.salary is null UNION SELECT Salaries.employee_id FROM Salaries LEFT JOIN Employees on Employees.employee_id = Salaries.employee_id WHERE  Employees.name is null ORDER BY employee_id;")
	if err != nil {
		return
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

	fmt.Println(employees)
}
