package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
)

type Student struct {
	Username string  `json:"username"`
	Score    float64 `json:"score"`
	Grade    string  `json:"grade"`
	Sex      string  `json:"sex"`
}

var (
	sexs    string
	Xxslice map[string]Student = make(map[string]Student, 1000000)
)

func Caidan() {
	Xinx /*菜单功能列表*/ := `|=========================================================|
||| s/S:查询所有学生的信息	| c/C:新增学生信息	|||
||| i/i:修改学生信息		| q/Q:退出程序		|||
|=========================================================|`

	color.Set(color.FgMagenta, color.Bold)
	defer color.Unset()
	fmt.Println(Xinx)

	color.Set(color.FgRed, color.Bold)
	fmt.Print("请您选择需要的操作:")
}

func Scans(a string) {
	var (
		username string
		score    float64
		grade    string
		sex      string
		sum      Student
	)
	switch a {
	case "s", "S":
		var a string
		fmt.Print("请输入您要查询的学生姓名,如需查询全部学生信息,请按l/L:")
		fmt.Scanf("%s\n", &a)
		if a == "l" || a == "L" {
			for _, value := range Xxslice {
				date, _ := json.Marshal(value)
				fmt.Printf("%v\n", date)
			}
		} else {
			_, ok := Xxslice[a]
			if ok == true {
				fmt.Printf("%v\n", Xxslice[a])
			} else {
				color.Set(color.FgRed, color.Bold)
				defer color.Unset()
				fmt.Println("您查询的学生不存在!")
			}
			fmt.Println()
		}
	case "c", "C":
		for i := 0; i < 4; i++ {
			fmt.Print("请输入新增学生的名字:")
			if i == 0 {
				fmt.Printf("请输入学生姓名:")
				fmt.Scanf("%s\n", &username)
				sum.Username = username
			} else if i == 1 {
				fmt.Printf("请输入学生班级:")
				fmt.Scanf("%s\n", &grade)
				sum.Grade = grade
			} else if i == 2 {
				fmt.Printf("请输入学生性别:")
				fmt.Scanf("%s\n", &sex)
				sum.Sex = sex
			} else if i == 3 {
				fmt.Printf("请输入学生分数:")
				fmt.Scanf("%f\n", &score)
				sum.Score = score
			}
		}
		Xxslice[username] = sum

	case "i", "I":
		var (
			num     int
			selects Student
		)
		for {
			fmt.Print("请选择您要修改的学生姓名：")
			fmt.Scanf("%v\n", &username)
			fmt.Print("请选择您要修改的学生信息（1：班级  2：性别  3：分数  4：全部）：")
			fmt.Scanf("%v\n", &num)
			if num == 1 {
				fmt.Printf("请输入%v的班级:", username)
				fmt.Scanf("%s\n", &grade)
				Xxslice[username] = Student{Grade: grade}

			} else if num == 2 {

			} else if num == 3 {

			} else if num == 4 {

			}
		}
	case "q", "Q":
		os.Exit(0)
	}

}

func main() {
	for {
		Caidan()
		fmt.Scan(&sexs)
		Scans(sexs)
	}
}
