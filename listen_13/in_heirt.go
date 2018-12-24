package main

import "fmt"

type Anima struct {
	Name string
	Sex  string
}

func (a *Anima) Talk() {
	fmt.Printf("%v\n", a.Name)
}

type Dog struct {
	Feet string
	*Anima
}

func (d Dog) Eat() {
	fmt.Println("Dog is Eat")
}

func main() {
	var b *Dog = &Dog{
		Feet: "foutf feet",
		Anima: &Anima{
			Name: "dog",
			Sex:  "xiong",
		},
	}
	b.Talk()
	b.Eat()
}
