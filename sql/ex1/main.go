package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pattarapong-ra/Hobbit/tree/main/exercise2/sql/ex2"
	"github.com/pattarapong-ra/Hobbit/tree/main/exercise2/sql/ex3"
)

type Employee struct {
	EmployeeID int
	Name       string
	Salary     int
}

type EmployeeQuery struct {
	EmployeeID int
	Bonus      int
}

// GetEmployees : get data all employees
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()

	rows, err := db.Query("SELECT * FROM employees;")
	if err != nil {
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.EmployeeID, &emp.Name, &emp.Salary); err != nil {
			fmt.Println(err)
		}
		employees = append(employees, emp)
	}

	fmt.Println(employees)

}

// 1. (Select) Write an SQL query to calculate the bonus of each employee.
// 		The bonus of an employee is 100% of their salary
//  	if the ID of the employee is an odd number and the employee name does not start with the character 'M'.
//  	The bonus of an employee is 0 otherwise.
// Return the result table ordered by employee_id.

// The query result format is in the following example.
// +-------------+-------+
// | employee_id | bonus |
// +-------------+-------+
// | 2           | 0     |
// | 3           | 0     |
// | 7           | 7400  |
// | 8           | 0     |
// | 9           | 7700  |
// +-------------+-------+

// CalculateBonus :  Calculate Special Bonus
func CalculateBonus(w http.ResponseWriter, r *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()

	rows, err := db.Query("SELECT * FROM employees;")
	if err != nil {
		return
	}
	defer rows.Close()

	var employeeBonus []EmployeeQuery
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
	}
	fmt.Println(employeeBonus)

}

// 2. (Update) Write an SQL update to fix the names so that only the first character is uppercase and the rest are lowercase.
// Return the result table ordered by employee_id.
// Employees table:
// +-------------+---------+
// | employee_id | name    |
// +-------------+---------+
// | 2           | Meir    |
// | 3           | Michael |
// | 7           | Addilyn |
// | 8           | Juan    |
// | 9           | Kannon  |
// +-------------+---------+
// FixNames : Fix Names in a Table
func FixNames(w http.ResponseWriter, r *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()
	_, err = db.Query("UPDATE Employees SET name=INITCAP(name);")
	if err != nil {
		return
	}

	rows, err2 := db.Query("SELECT employee_id,name, salary FROM Employees;")
	if err2 != nil {
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.EmployeeID, &emp.Name, &emp.Salary); err != nil {
			fmt.Println(err)
		}
		employees = append(employees, emp)
	}
	fmt.Println(employees)

}

// 3. (Delete) Write delete SQL for employees with bonuses.
// Return the result table ordered by employee_id.
// Employees table:
// +-------------+---------+
// | employee_id | name    |
// +-------------+---------+
// | 2           | Meir    |
// | 3           | Michael |
// | 8           | Juan    |
// +-------------+---------+
// RemoveSomeEmployee :
func RemoveSomeEmployee(w http.ResponseWriter, r *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()

	_, err = db.Query("DELETE FROM Employees WHERE (NOT name LIKE 'M%') AND mod(employee_id,2) <> 0;")
	if err != nil {
		return
	}

	rows, err2 := db.Query("SELECT * FROM Employees;")
	if err2 != nil {
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.EmployeeID, &emp.Name, &emp.Salary); err != nil {
			fmt.Println(err)
		}
		employees = append(employees, emp)
	}
	fmt.Println(employees)

}

const (
	host     = "localhost"
	port     = 5432
	user     = "mooham12314"
	password = "mooham12315"
	dbname   = "postgres"
)

func InitializeDB() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()
	defer fmt.Println("Closed!")
	// check db
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected!")
	//PrepareTable1(db)
}

func main() {

	InitializeDB()
	r := mux.NewRouter()
	r.HandleFunc("/get", GetEmployees).Methods("GET")
	r.HandleFunc("/get/calBonus", CalculateBonus).Methods("GET")
	r.HandleFunc("/get/fixNames", FixNames).Methods("GET")
	r.HandleFunc("/get/remove-employees", RemoveSomeEmployee).Methods("GET")
	r.HandleFunc("/get/employees-missing-information", ex2.GetEmployeesWithMissingInformation).Methods("GET")
	r.HandleFunc("/get/swapSalary", ex3.SwapSalary).Methods("GET")
	log.Fatal(http.ListenAndServe(":8002", r))

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
