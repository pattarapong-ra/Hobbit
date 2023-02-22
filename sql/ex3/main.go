package ex3

import (
	"net/http"
)

// Write an SQL query to swap all 'f' and 'm' values (i.e., change all 'f' values to 'm' and vice versa)
// with a single update statement and no intermediate temporary tables.
// Note that you must write a single update statement, do not write any select statement for this problem.

// The query result format is in the following example.
// Output:
// +----+------+-----+--------+
// | id | name | sex | salary |
// +----+------+-----+--------+
// | 1  | A    | f   | 2500   |
// | 2  | B    | m   | 1500   |
// | 3  | C    | f   | 5500   |
// | 4  | D    | m   | 500    |
// +----+------+-----+--------+

// SwapSalary :Swap Salary
func SwapSalary(w http.ResponseWriter, r *http.Request) {

}
