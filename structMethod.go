package main

import (
	"fmt"
)

type Persona struct {
	Id   int
	Name string
}

// UpdateName не изменит оригинальной структуры,
//для которой вызван
func (p Persona) UpdateName(name string) {
	p.Name = name
}

// SetName изменяет оригинальную структуру
func (p *Persona) SetName(name string) {
	p.Name = name
}

type Accounta struct {
	Id   int
	Name string
	Persona
}

func (p *Accounta) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) Add(in int) {
	*sl = append(*sl, in)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func StructMethod() {
	var pers = Persona{Id: 1}
	fmt.Println(pers)
	pers.SetName("NewName")
	//(&pers).SetName("Hello")
	fmt.Println(pers)

	var acc = Accounta{1, "rvasily", Persona{2, "Vasily Romanov"}}
	acc.SetName("Name")
	acc.Persona.SetName("Name2")
	fmt.Printf("%#v \n", acc)

	sl := MySlice([]int{1, 2})
	sl.Add(5)
	fmt.Println(sl.Count(), sl)
}
