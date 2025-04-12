package main

import (
	"final/tasks"
	"fmt"
)

func main() {
	//Task 1
	fmt.Println(tasks.Polindrome("scooter"))
	fmt.Println(tasks.Polindrome("race   ,,,  ... car"))
	//Task 2
	e := tasks.NewEmployee("John", "Admin", 50000)
	salary := e.UpSalary(20)
	fmt.Println("Новая зарплата сотрудника ", salary)
	//Task 3
	slice := []int{3, 6, 8, 2, 4}
	fmt.Println(tasks.SortSlice(slice))
	//Task 4
	slice2 := []int{3, 6, 8, 2, 4}
	fmt.Println(tasks.EqualSlices(slice, slice2))
	//Task 5
	fmt.Println(tasks.GetLetters("gdh8hfhf9hdf0fb945"))
	//Task 6
	tasks.ReadSqrt()
	//Task 7
	stack := tasks.NewStack([]int{1, 2, 3, 4, 5})
	stack.Print()
	stack.Push(7)
	stack.Print()
	stack.Pop()
	stack.Print()
	fmt.Println(stack.Top())
	fmt.Println(stack.IsEmpty())
}
