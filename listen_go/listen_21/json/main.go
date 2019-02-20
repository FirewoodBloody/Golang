package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func WriteJson(filename string) (err error) {
	_ = filename
	var person []*Person
	for i := 0; i < 10; i++ {
		p := &Person{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(int(time.Now().Unix())),
			Sex:  "男",
		}
		person = append(person, p)
	}
	data, err := json.Marshal(person)
	fmt.Printf("%s", data)
	//if err != nil {
	//	fmt.Printf("marshal failed,err:%s\n", err)
	//	return
	//}
	//err = ioutil.WriteFile(filename, data, 755)
	//if err != nil {
	//	fmt.Printf("write failed,err:%s\n", err)
	//	return
	//}
	return
}

func ReadJson(filename string) (err error) {
	var person []*Person
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Read file failed,err:%s\n", err)
		return
	}
	err = json.Unmarshal(data, &person)
	if err != nil {
		fmt.Printf("unmarshal failed,err:%s\n", err)
		return
	}
	for _, v := range person {
		fmt.Printf("%#v\n", v)
	}
	return
}

func main() {
	filename := "D:/logs/json.txt"
	err := WriteJson(filename)
	if err != nil {
		fmt.Printf("Write json failed,err:%s\n", err)
		return
	}
	//
	//err = ReadJson(filename)
	//if err != nil {
	//	fmt.Printf("Read json failed,err:%s\n", err)
	//	return
	//}
}
