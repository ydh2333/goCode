package main


func main2() {
	//数组定义1，先定义后赋值
	var arr1 [5] int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50

	for i,v := range arr1{
		println("arr1:",i,v)
	}
	println("----------------------")
	//数组定义2，声明时初始化
	arr2 := [5]int{11,22,33,44,55}

	for i,v := range arr2{
		println("arr2:",i,v)
	}

	println("----------------------")
	//数组定义3，使用...代替数组长度
	arr3 := [...]int{5,10,15,20,25,30}
	for i,v := range arr3{
		println("arr3:",i,v)
	}

	println("----------------------")
	//数组定义4，指定数组长度，初始化元素
	arr4 := [5]int{99,88,77,66,55}
	for i,v := range arr4{
		println("arr4:",i,v)
	}
	arr5 := [5]int{0:10,2:20,4:22}
	for i,v := range arr5{
		println("arr5:",i,v)
	}





}