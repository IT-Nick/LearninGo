package main

import "fmt"

func Variables() {
	//переменные
	var num0 int
	var num1 int = 1
	var num2 = 20
	fmt.Println(num0, num1, num2)

	//коротко так
	num := 30
	fmt.Println(num)

	var weight, height int = 10, 20

	weight, height = 11, 21

	weight, age := 12, 22

	fmt.Println(weight, height, age)

	var i int = 10

	var autoInt = -10

	var bigInt int64 = 1<<32 - 1

	var unsignedInt uint = 100500

	fmt.Println(i, autoInt, bigInt, unsignedInt)

}
