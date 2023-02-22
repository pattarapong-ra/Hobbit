package ex1

import (
	"log"
	"net/http"

	"github.com/go-beginner-training/sql/ex1"
	"github.com/go-beginner-training/sql/ex2"
	"github.com/go-beginner-training/sql/ex3"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get", ex1.GetEmployees).Methods("GET")
	r.HandleFunc("/get/calBonus", ex1.CalculateBonus).Methods("GET")
	r.HandleFunc("/get/fixNames", ex1.FixNames).Methods("GET")
	r.HandleFunc("/get/remove-employees", ex1.RemoveSomeEmployee).Methods("GET")
	r.HandleFunc("/get/employees-missing-information", ex2.GetEmployeesWithMissingInformation).Methods("GET")
	r.HandleFunc("/get/swapSalary", ex3.SwapSalary).Methods("GET")
	log.Fatal(http.ListenAndServe(":8002", r))
}
