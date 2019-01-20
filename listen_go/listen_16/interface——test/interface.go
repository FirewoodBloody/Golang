package main

import "fmt"

type Ainmal interface {
	Talk()
	Eat()
	Name() string
}

type Dog struct {
}

func (d Dog) Talk() {
	fmt.Println("汪汪汪")
}

func (d Dog) Eat() {
	fmt.Println("我在吃骨头")
}

func (d Dog) Name() string {
	fmt.Println("我的名字叫旺财")
	return "旺财"
}

type Pag struct {
}

func (p Pag) Talk() {
	fmt.Println("汪汪汪")
}

func (p Pag) Eat() {
	fmt.Println("我在吃骨头")
}

func (p Pag) Name() string {
	fmt.Println("我的名字叫旺财")
	return "旺财"
}

func main() {
	var (
		b Dog
		a Ainmal
	)
	a = b
	a.Name()
	a.Eat()
	a.Talk()

	var p Pag
	a = p
	a.Name()
	a.Eat()
	a.Talk()
}
