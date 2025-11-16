package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Println("姓名：", e.Person.Name, "，年龄：", e.Person.Age, "，员工ID：", e.EmployeeID)
}

func main6() {
	e := Employee{
		EmployeeID: "9527",
		Person: Person{
			Name: "张三",
			Age: 25,
		},
	}
	e.PrintInfo()
}