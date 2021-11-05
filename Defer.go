package main

import "fmt"

func getSomeVars() string {
	fmt.Println("Аргументы функции")
	return "первый дефер через getSomeVars()"
}

func Defer() {
	//Defer - отложенное выполнение функции
	//Это когда какая-то работа будет
	//выполнена после завершения функции.
	defer fmt.Println("второй дефер")
	//тут вычисления функции (не то что она
	//вернет) выполнятся при объявлении defer
	//а потом уже при выполнении defer
	//выполнится результат функции
	defer fmt.Println(getSomeVars())

	//но можно через анонимную функцию
	//тогда все что в функции будет сразу
	//просчитано
	defer func() {
		fmt.Println(getSomeVars())
	}()

	fmt.Println("Обычный fmt")
}
