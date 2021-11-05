package main

import "fmt"

func Defer() {
	//Defer - отложенное выполнение функции
	//Это когда какая-то работа будет
	//выполнена после завершения функции.


	defer fmt.Println("After work")
	fmt.Println("Some userful work")
}
