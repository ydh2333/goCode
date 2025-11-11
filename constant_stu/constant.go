package main

import "fmt"

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

// func (g Gender) CaseGen() string {
// 	fmt.Println(g)
// 	fmt.Println(&g)
// 	switch g {
// 	case Male:
// 		return "Male"
// 	case Female:
// 		return "Female"
// 	default:
// 		return "Unknown"
// 	}
// }

func (g *Gender) ABC() string {
	fmt.Printf("g的地址是：%p\n", g)
	switch *g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown"
	}
}

func main() {
	var g Gender = "Female"
	fmt.Println("g:", g)
	fmt.Printf("g的地址是：%p\n", &g)

	fmt.Println("-------------------------------")

	// switch g {
	// case Male:
	// 	fmt.Println("结果是：", g)
	// case Female:
	// 	fmt.Println("结果是：", g)
	// default:
	// 	fmt.Println("未知")
	// }
	fmt.Println("结果是：", g.ABC())

	fmt.Println("-------------------------------")

	var d interface{}
	d = "130"

	switch t := d.(type) {
	case byte:
		fmt.Println("d is byte type, ", t)
	case *byte:
		fmt.Println("d is byte point type, ", t)
	case *int:
		fmt.Println("d is int type, ", t)
	case *string:
		fmt.Println("d is string type, ", t)
	default:
		fmt.Println("d is unknown type, ", t)
	}

	fmt.Println("-------------------------------")
	
	var b int = 4
	fmt.Println("local variable, b = ", b)
	if b := 3; b == 3 {
		fmt.Println("if statement, b = ", b)
		b--
		fmt.Println("if statement, b = ", b)
	}
	fmt.Println("local variable, b = ", b)
}
