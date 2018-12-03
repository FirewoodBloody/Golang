package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

type Student struct {
	Username string  `json:"username"`
	Score    float64 `json:"score"`
	Grade    string  `json:"grade"`
	Sex      string  `json:"sex"`
}

var (
	Xxslice map[string]Student = make(map[string]Student, 1000000)
)

func Caidan() {
	Xinx /*菜单功能列表*/ := `|=========================================================|
|||			操作菜单			|||
|=========================================================|
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
		num      int
		yes      string
	)
	if a != "" {
		switch a {
		case "s", "S":
			var b string
			fmt.Print("请输入您要查询的学生姓名,如需查询全部学生信息,请按l/L:")
			fmt.Scanf("%s\n", &b)
			if b == "l" || b == "L" {
				if len(Xxslice) != 0 {
					for _, value := range Xxslice {
						date, _ := json.Marshal(value)
						fmt.Printf("%#s\n", string(date))
					}
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
					break
				} else {
					color.Set(color.FgRed, color.Bold)
					defer color.Unset()
					fmt.Println()
					time.Sleep(time.Second * 1)
					fmt.Println("当前无任何学生信息！")
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
					break
				}
			} else {
				_, ok := Xxslice[b]
				if ok == true {
					time.Sleep(time.Second * 1)
					fmt.Println()
					date, _ := json.Marshal(Xxslice[b])
					fmt.Printf("%v\n", string(date))
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
					break
				} else {
					color.Set(color.FgRed, color.Bold)
					defer color.Unset()
					time.Sleep(time.Second * 1)
					fmt.Println()
					fmt.Println("您查询的学生不存在!")
					time.Sleep(time.Second * 2)
					fmt.Println("\n")

					break
				}
				fmt.Println()
				break
			}
		case "c", "C":

			fmt.Printf("请输入学生姓名:")
			fmt.Scanf("%s\n", &username)
			sum.Username = username

			fmt.Printf("请输入学生班级:")
			fmt.Scanf("%s\n", &grade)
			sum.Grade = grade

			fmt.Printf("请输入学生性别:")
			fmt.Scanf("%s\n", &sex)
			sum.Sex = sex

			fmt.Printf("请输入学生分数:")
			fmt.Scanf("%f\n", &score)
			sum.Score = score

			fmt.Printf("请确认是否创建%v的信息！（y/n）:", sum)
			fmt.Scanf("%s\n", &yes)
			if yes == "y" || yes == "Y" {
				Xxslice[username] = sum
				fmt.Printf("%v 信息创建成功！\n", Xxslice[username])
				time.Sleep(time.Second * 2)
				fmt.Println("\n")
				break
			} else {
				fmt.Println("信息创建已取消！")
				time.Sleep(time.Second * 2)
				fmt.Println("\n")
				break
			}
		case "i", "I":

			fmt.Print("请选择您要修改的学生姓名：")
			fmt.Scanf("%v\n", &username)
			fmt.Print("请选择您要修改的学生信息（1：班级  2：性别  3：分数  4：全部）：")
			fmt.Scanf("%v\n", &num)
			if num == 1 {
				fmt.Printf("请输入%v的班级:", username)
				fmt.Scanf("%s\n", &grade)
				fmt.Print("请确认将%v的班级修改为%v (y/n):", username, grade)
				fmt.Scanf("%s\n", yes)
				if yes == "y" || yes == "Y" {
					Xxslice[username] = Student{Grade: grade}
					color.Set(color.FgGreen, color.Bold)
					defer color.Unset()
					fmt.Printf("学生%v的班级信息修改成功！\n", Xxslice[username])
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
					break
				} else {
					fmt.Println("信息修改已取消！")
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
					break
				}
			} else if num == 2 {
				fmt.Printf("请输入%v的性别:", username)
				fmt.Scanf("%s\n", &sex)
				fmt.Print("请确认将%v的性别修改为%v (y/n):", username, sex)
				fmt.Scanf("%s\n", yes)
				if yes == "y" || yes == "Y" {
					Xxslice[username] = Student{Sex: sex}
					color.Set(color.FgGreen, color.Bold)
					defer color.Unset()
					fmt.Printf("学生%v的性别信息修改成功！\n", Xxslice[username])
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
					break
				} else {
					fmt.Println("信息修改已取消！")
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
					break
				}
			} else if num == 3 {
				fmt.Printf("请输入%v的分数:", username)
				fmt.Scanf("%s\n", &score)
				fmt.Print("请确认将%v的分数修改为%v (y/n):", username, score)
				fmt.Scanf("%s\n", yes)
				if yes == "y" || yes == "Y" {
					Xxslice[username] = Student{Score: score}
					color.Set(color.FgGreen, color.Bold)
					defer color.Unset()
					fmt.Printf("学生%v的性别信息修改成功！\n", Xxslice[username])
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
					break
				} else {
					fmt.Println("信息修改已取消！")
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
					break
				}
			} else if num == 4 {
				fmt.Printf("请输入学生班级:")
				fmt.Scanf("%s\n", &grade)

				fmt.Printf("请输入学生性别:")
				fmt.Scanf("%s\n", &sex)

				fmt.Printf("请输入学生分数:")
				fmt.Scanf("%f\n", &score)

				fmt.Print("请确认将%v的信息修改为:班级：%v  性别：%v  分数：%v (y/n):", username, grade, sex, score)
				fmt.Scanf("%v\n", &yes)
				if yes == "y" || yes == "Y" {
					Xxslice[username] = Student{
						Grade: grade,
						Sex:   sex,
						Score: score,
					}
					color.Set(color.FgGreen, color.Bold)
					defer color.Unset()
					fmt.Printf("学生%v的性别信息修改成功！\n", Xxslice[username])
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
				} else {
					fmt.Println("信息修改已取消！")
					time.Sleep(time.Second * 2)
					fmt.Println("\n")
					break
				}
			}

		case "q", "Q":
			os.Exit(0)
		}

	} else {
		color.Set(color.FgRed, color.Bold)
		defer color.Unset()
		fmt.Println("您未进行选择，请您选择需要的操作！")
		time.Sleep(time.Second * 1)
		fmt.Println("\n")
	}
}

func main() {
	for {
		var sexs string
		Caidan()
		fmt.Scanf("%s\n", &sexs)
		Scans(sexs)
	}
}
