package tasks

// Создайте структуру "Сотрудник" с полями (имя, должность, зарплата). Реализуйте метод для увеличения зарплаты сотрудника на заданный процент.

type Employee struct {
	name   string
	post   string
	salary int
}

func NewEmployee(name string, post string, salary int) Employee {
	return Employee{
		name:   name,
		post:   post,
		salary: salary,
	}
}

func (e Employee) UpSalary(up int) int {
	return e.salary * (100 + up) / 100
}
