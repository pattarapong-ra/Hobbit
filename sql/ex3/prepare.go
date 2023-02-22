package ex3

import "database/sql"

func PrepareTable(db *sql.DB) error {
	err := createtableSalary(db)
	if err != nil {
		return err
	}
	return nil
}
func createtableSalary(db *sql.DB) error {
	// 	Table: Salary
	// +-------------+----------+
	// | Column Name | Type     |
	// +-------------+----------+
	// | id          | int      |
	// | name        | varchar  |
	// | sex         | ENUM     |
	// | salary      | int      |
	// +-------------+----------+
	// id is the primary key for this table.
	// The sex column is ENUM value of type ('m', 'f').
	// The table contains information about an employee.

	dropBeforeCreateTable := `DROP TABLE IF EXISTS Salary ;`
	_, errdropBeforeCreateTable := db.Exec(dropBeforeCreateTable)
	if errdropBeforeCreateTable != nil {
		return errdropBeforeCreateTable
	}
	createTable := `Create table If Not Exists Salary (
			id int, 
			name varchar(100), 
			sex char(1), 
			salary int
		)
	;`
	_, errcreateTable := db.Exec(createTable)
	if errcreateTable != nil {
		return errcreateTable
	}
	// Salary table:
	// +----+------+-----+--------+
	// | id | name | sex | salary |
	// +----+------+-----+--------+
	// | 1  | A    | m   | 2500   |
	// | 2  | B    | f   | 1500   |
	// | 3  | C    | m   | 5500   |
	// | 4  | D    | f   | 500    |
	// +----+------+-----+--------+

	insert := `insert into Salary (id, name, sex, salary) values ('1', 'A', 'm', '2500')` +
		`insert into Salary (id, name, sex, salary) values ('2', 'B', 'f', '1500')` +
		`insert into Salary (id, name, sex, salary) values ('3', 'C', 'm', '5500')` +
		`insert into Salary (id, name, sex, salary) values ('4', 'D', 'f', '500')`

	_, errinsert := db.Exec(insert)
	if errinsert != nil {
		return errinsert
	}
	return nil

}
