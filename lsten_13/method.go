package main

import "fmt"

type Porelo struct {
	Name    string
	Country string
}

var a int

type Intfunc int

func (i *Intfunc) echo() {
	fmt.Println(*i)
	*i = 2
}

func (p Porelo) echo() {
	p.Country = "meigou"
	fmt.Println(p.Name, p.Country)
}

func main() {
	a := Intfunc(100)
	p1 := Porelo{
		Name:    "张三",
		Country: "中国",
	}
	p1.echo()
	fmt.Printf("%#v\n", p1)
	a.echo()
	fmt.Println(a)
}
