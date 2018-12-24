package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int
	Name string
	Sex  string
}

type Class struct {
	Name     string
	Count    int
	Students []*Student
}

var rawjson = `
{"Name":"101","Count":0,"Students":[{"Id":0,"Name":"stu0","Sex":"man"},{"Id":1,"Name":"stu1","Sex":"man"},{"Id":2,"Name":"stu2","Sex":"man"},{"Id":3,"Name":"stu3","Sex":"man"},{"Id":4,"Name":"stu4","Sex":"man"},{"Id":5,"Name":"stu5","Sex":"man"},{"Id":6,"Name":"stu6","Sex":"man"},{"Id":7,"Name":"stu7","Sex":"man"},{"Id":8,"Name":"stu8","Sex":"man"},{"Id":9,"Name":"stu9","Sex":"man"},{"Id":10,"Name":"stu10","Sex":"man"},{"Id":11,"Name":"stu11","Sex":"man"},{"Id":12,"Name":"stu12","Sex":"man"},{"Id":13,"Name":"stu13","Sex":"man"},{"Id":14,"Name":"stu14","Sex":"man"},{"Id":15,"Name":"stu15","Sex":"man"},{"Id":16,"Name":"stu16","Sex":"man"},{"Id":17,"Name":"stu17","Sex":"man"},{"Id":18,"Name":"stu18","Sex":"man"},{"Id":19,"Name":"stu19","Sex":"man"},{"Id":20,"Name":"stu20","Sex":"man"},{"Id":21,"Name":"stu21","Sex":"man"},{"Id":22,"Name":"stu22","Sex":"man"},{"Id":23,"Name":"stu23","Sex":"man"},{"Id":24,"Name":"stu24","Sex":"man"},{"Id":25,"Name":"stu25","Sex":"man"},{"Id":26,"Name":"stu26","Sex":"man"},{"Id":27,"Name":"stu27","Sex":"man"},{"Id":28,"Name":"stu28","Sex":"man"},{"Id":29,"Name":"stu29","Sex":"man"},{"Id":30,"Name":"stu30","Sex":"man"},{"Id":31,"Name":"stu31","Sex":"man"},{"Id":32,"Name":"stu32","Sex":"man"},{"Id":33,"Name":"stu33","Sex":"man"},{"Id":34,"Name":"stu34","Sex":"man"},{"Id":35,"Name":"stu35","Sex":"man"},{"Id":36,"Name":"stu36","Sex":"man"},{"Id":37,"Name":"stu37","Sex":"man"},{"Id":38,"Name":"stu38","Sex":"man"},{"Id":39,"Name":"stu39","Sex":"man"},{"Id":40,"Name":"stu40","Sex":"man"},{"Id":41,"Name":"stu41","Sex":"man"},{"Id":42,"Name":"stu42","Sex":"man"},{"Id":43,"Name":"stu43","Sex":"man"},{"Id":44,"Name":"stu44","Sex":"man"},{"Id":45,"Name":"stu45","Sex":"man"},{"Id":46,"Name":"stu46","Sex":"man"},{"Id":47,"Name":"stu47","Sex":"man"},{"Id":48,"Name":"stu48","Sex":"man"},{"Id":49,"Name":"stu49","Sex":"man"},{"Id":50,"Name":"stu50","Sex":"man"},{"Id":51,"Name":"stu51","Sex":"man"},{"Id":52,"Name":"stu52","Sex":"man"},{"Id":53,"Name":"stu53","Sex":"man"},{"Id":54,"Name":"stu54","Sex":"man"},{"Id":55,"Name":"stu55","Sex":"man"},{"Id":56,"Name":"stu56","Sex":"man"},{"Id":57,"Name":"stu57","Sex":"man"},{"Id":58,"Name":"stu58","Sex":"man"},{"Id":59,"Name":"stu59","Sex":"man"},{"Id":60,"Name":"stu60","Sex":"man"},{"Id":61,"Name":"stu61","Sex":"man"},{"Id":62,"Name":"stu62","Sex":"man"},{"Id":63,"Name":"stu63","Sex":"man"},{"Id":64,"Name":"stu64","Sex":"man"},{"Id":65,"Name":"stu65","Sex":"man"},{"Id":66,"Name":"stu66","Sex":"man"},{"Id":67,"Name":"stu67","Sex":"man"},{"Id":68,"Name":"stu68","Sex":"man"},{"Id":69,"Name":"stu69","Sex":"man"},{"Id":70,"Name":"stu70","Sex":"man"},{"Id":71,"Name":"stu71","Sex":"man"},{"Id":72,"Name":"stu72","Sex":"man"},{"Id":73,"Name":"stu73","Sex":"man"},{"Id":74,"Name":"stu74","Sex":"man"},{"Id":75,"Name":"stu75","Sex":"man"},{"Id":76,"Name":"stu76","Sex":"man"},{"Id":77,"Name":"stu77","Sex":"man"},{"Id":78,"Name":"stu78","Sex":"man"},{"Id":79,"Name":"stu79","Sex":"man"},{"Id":80,"Name":"stu80","Sex":"man"},{"Id":81,"Name":"stu81","Sex":"man"},{"Id":82,"Name":"stu82","Sex":"man"},{"Id":83,"Name":"stu83","Sex":"man"},{"Id":84,"Name":"stu84","Sex":"man"},{"Id":85,"Name":"stu85","Sex":"man"},{"Id":86,"Name":"stu86","Sex":"man"},{"Id":87,"Name":"stu87","Sex":"man"},{"Id":88,"Name":"stu88","Sex":"man"},{"Id":89,"Name":"stu89","Sex":"man"},{"Id":90,"Name":"stu90","Sex":"man"},{"Id":91,"Name":"stu91","Sex":"man"},{"Id":92,"Name":"stu92","Sex":"man"},{"Id":93,"Name":"stu93","Sex":"man"},{"Id":94,"Name":"stu94","Sex":"man"},{"Id":95,"Name":"stu95","Sex":"man"},{"Id":96,"Name":"stu96","Sex":"man"},{"Id":97,"Name":"stu97","Sex":"man"},{"Id":98,"Name":"stu98","Sex":"man"},{"Id":99,"Name":"stu99","Sex":"man"}]}
`

func main() {
	c := &Class{
		Name:  "101",
		Count: 0,
	}
	for i := 0; i < 100; i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%d", i),
			Id:   i,
			Sex:  "man",
		}
		c.Students = append(c.Students, stu)
	}
	value, err := json.Marshal(c)
	if err == nil {
		fmt.Printf("json:%s\n", string(value))
	} else {
		fmt.Println("json maeshal err")
	}

	var c1 *Class = &Class{}
	err = json.Unmarshal([]byte(rawjson), c1)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Printf("%#v", c1)
}
