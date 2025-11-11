package main

import "fmt"

func main2() {

	type Student struct {
		Name  string
		Age   int
		Score float64
	}

	var stu1 Student = Student{
		Name:  "Alice",
		Age:   20,
		Score: 88.5,
	}

	fmt.Println("name:", stu1.Name, "Age:", stu1.Age, "Score:", stu1.Score)

	fmt.Println("------------------")

	type Student2 struct {
		string
		int
		float64
	}

	var stu2 Student2 = Student2{
		"tom",
		18,
		99.5,
	}

	fmt.Println("name:", stu2.string, "Age:", stu2.int, "Score:", stu2.float64)

	stu2.int = 22

	fmt.Println("name:", stu2.string, "Age:", stu2.int, "Score:", stu2.float64)

	fmt.Println("------------------")

	var stu3 = struct {
		Name  string
		Age   int
		Score float64
	}{
		Name:  "jack",
		Age:   19,
		Score: 77.5,
	}

	fmt.Println("str3:", stu3, "Name:", stu3.Name)




}