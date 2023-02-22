package ex2

import (
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

// getEmployeesWithMissingInformation : Employees With Missing Information
func GetEmployeesWithMissingInformation(w http.ResponseWriter, r *http.Request) {

}
