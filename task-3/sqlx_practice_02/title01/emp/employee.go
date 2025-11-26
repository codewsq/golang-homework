package emp

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

/*
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

/*
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中
*/
func EmpSelectByDep(db *sqlx.DB) {
	sqlStr := "select * from employees where department = ?"
	rows, err := db.Query(sqlStr, "技术部")
	if err != nil {
		panic(err)
	}
	// 最终结束后释放连接
	defer rows.Close()

	var emps []Employee

	for rows.Next() {
		var emp Employee
		err = rows.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Salary)
		emps = append(emps, emp)
	}

	fmt.Println("技术部员工数量为：", len(emps))
	for _, emp := range emps {
		fmt.Println("技术部员工信息：ID:", emp.ID, ",Name:", emp.Name, ",department:", emp.Department, ",salary:", emp.Salary)
	}
}

/*
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
func EmpMaxSalary(db *sqlx.DB) {
	sqlStr := "select * from employees order by salary desc limit 1"
	var emp Employee
	err := db.Get(&emp, sqlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("工资最高的员工信息：ID:", emp.ID, ",Name:", emp.Name, ",department:", emp.Department, ",salary:", emp.Salary)
}
