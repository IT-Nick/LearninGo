package main

import "fmt"

func deferTest() {
	//отловка паники, чтобы другая часть программы не крашнулась
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend:", err)
		}
	}()
	fmt.Println("Some useful work")
	panic("something bad happend")
}

func panicDefer() {
	//бывает что и в defer мы упали в панику
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend FIRST:", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend SECOND:", err)
			panic("second panic")
		}
	}()
	fmt.Println("Some usefil work")
	panic("something bad happend too")
}

func Panics() {
	deferTest()
	panicDefer()
	return
}
