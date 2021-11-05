package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person //в данном случае это свойство структуры
	Person         //это уже часть самой структуры
}

func Struct() {
	//полный формат объявления
	var acc = Account{Id: 1, Name: "stepan", Person: Person{
		Name:    "Alex",
		Address: "Москва",
	}}

	fmt.Printf("%#v\n", acc)

	//короткая форма
	acc.Owner = Person{2, "Stepan", "Moscow"}
	fmt.Printf("%#v\n", acc)

	//можем напрямую работать с Person через acc потому что мы встроили
	//Person в аккаунт
	acc.Id = 3
	acc.Address = "Москаль"
	fmt.Println(acc.Id, acc.Address, acc.Person.Name, acc.Name)

}
