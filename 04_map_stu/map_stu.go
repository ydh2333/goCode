package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println("m1 =", m1)

	fmt.Println("---------------------------------------")

	var m2 map[string]string = map[string]string{
		"one":   "hello",
		"two":   "world",
		"three": "golang",
	}
	m2["four"] = "java"
	fmt.Println("m2 =", m2, "\nm2[\"one\"]", m2["one"])

	fmt.Println("---------------------------------------")

	m2One, exist := m2["one1"]
	if exist {
		fmt.Println("m2One:", m2One)
	} else {
		fmt.Println("m2One does not exist")
	}

	fmt.Println("---------------------------------------")

	for k, v := range m2 {
		fmt.Println("key:", k, "value:", v)
	}

	fmt.Println("---------------------------------------")

	m3 := make(map[int]string)
	m3[1] = "apple"
	fmt.Println("m3 =", m3, "len(m3) =", len(m3))

	fmt.Println("---------------------------------------")

	delete(m2, "one")

	fmt.Println(m2)

}
