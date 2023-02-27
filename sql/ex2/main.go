package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pattarapong-ra/Hobbit/tree/main/exercise2/sql/ex3"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mooham12314"
	password = "mooham12315"
	dbname   = "postgres"
)

type EmployeeName struct {
	EmployeeID int
	Name       string
}

type EmployeeSalary struct {
	EmployeeID int
	Salary     int
}

// GetEmployees : get data all employees
func GetEmployeeName(w http.ResponseWriter, r *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Employees;")
	if err != nil {
		return
	}
	defer rows.Close()

	var employees []EmployeeName
	for rows.Next() {
		var emp EmployeeName
		if err := rows.Scan(&emp.EmployeeID, &emp.Name); err != nil {
			fmt.Println(err)
		}
		employees = append(employees, emp)
	}

	fmt.Println(employees)

}
func GetEmployeeSalary(w http.ResponseWriter, r *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Salaries;")
	if err != nil {
		return
	}
	defer rows.Close()

	var employees []EmployeeSalary
	for rows.Next() {
		var emp EmployeeSalary
		if err := rows.Scan(&emp.EmployeeID, &emp.Salary); err != nil {
			fmt.Println(err)
		}
		employees = append(employees, emp)
	}

	fmt.Println(employees)

}

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
	PrepareTable2(db)
}

func main() {
	InitializeDB()
	r := mux.NewRouter()
	r.HandleFunc("/getname", GetEmployeeName).Methods("GET")
	r.HandleFunc("/getsalary", GetEmployeeSalary).Methods("GET")
	r.HandleFunc("/get/employees-missing-information", GetEmployeesWithMissingInformation).Methods("GET")
	r.HandleFunc("/get/swapSalary", ex3.SwapSalary).Methods("GET")
	log.Fatal(http.ListenAndServe(":8002", r))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
