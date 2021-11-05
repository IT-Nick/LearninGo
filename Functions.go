package main

import "fmt"
//обычная функция
func doNothing(){
	fmt.Println("I do nothing")
}

func Functions() {
	doNothing()

	//Анонимная функция
	func(in string){
		fmt.Println("anon func say:", in)
	}("nothing")//сразу можно передать nothing


	//присвоим анонимн функ к переменной
	printer := func(in string) {
		fmt.Println("printer =", in)
	}

	printer("as variable")

	//создадим кастомный тип "сигратура функции со стрингом"
	type strFuncType func(string)

	//используем наш кастомный тип сигнатуры функции
	//чтобы передать целую функцию! в нашу новую функцию
	worker := func(callback strFuncType) {
		callback("as callback")
	}
	worker(printer)

	//функция возвращает замыкание
	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}

	successLogger := prefixer("SUCCESS")
	successLogger("expected behaviour")

}
