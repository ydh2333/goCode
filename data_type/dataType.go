package main

import "fmt"

func main() {
	var a uint8 = 0xe
	fmt.Println("a unit8:", a) 
	
	fmt.Println("-----------------------------------")

	var s1 string = "hello golang"
	fmt.Println("s string:", s1)

	fmt.Println("-----------------------------------")

	var bytes []byte = []byte(s1)
	fmt.Println("bytes slice:", bytes)

	fmt.Println("-----------------------------------")

	var s string = "hello Go"
	var bytes2 []byte = []byte(s)
	var runes []rune = []rune(s)

	fmt.Println("string length: ", len(s))
	fmt.Println("bytes length: ", len(bytes2))
	fmt.Println("runes length: ", len(runes))
}