package iniconfig

import (
	"fmt"
	"github.com/fatih/color"
	"reflect"
	"strings"
)

var listconfname string

func main() {

}
func Marshal(data interface{}) (result []byte, err error) {
	return
}

func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() == reflect.Ptr {
		if typeInfo.Elem().Kind() == reflect.Struct {
			for index, value := range lineArr {
				value = strings.TrimSpace(value)
				if len(value) == 0 || value[0] == ';' || value[0] == '#' {
					continue
				}
				if value[0] == '[' && value[len(value)-1] == ']' && len(value) > 2 {
					startName := strings.TrimSpace(value[1 : len(value)-1])
					if len(startName) != 0 {
						for i := 0; i < typeInfo.Elem().NumField(); i++ {
							if startName == typeInfo.Elem().Field(i).Tag.Get("ini") {
								listconfname = typeInfo.Elem().Field(i).Tag.Get("ini")
								fmt.Println(listconfname)
								break
							}
						}
					} else {
						//	err = fmt.Errorf("WARN: Skip line %d of the file and read the error", index+1)
						Color_fmt := color.New(color.FgYellow).Add(color.Underline)
						Color_fmt.Print("[WARN]: ")
						fmt.Printf("Skip line %d of the file and read the error\n", index+1)
						continue
					}

				} else if value[0] != '[' && value[len(value)-1] != ']' && len(value) > 2 {
					value = strings.TrimSpace(value)
					posiTion := strings.Index(value, "=")
					if posiTion == -1 || posiTion == 0 {
						Color_fmt := color.New(color.FgYellow).Add(color.Underline)
						Color_fmt.Print("[WARN]: ")
						fmt.Printf("Skip line %d of the file and read the error\n", index+1)
						continue
					} else {
						//key := strings.TrimSpace(value[0:posiTion])
						//val := strings.TrimSpace(value[posiTion:])
						resultValue := reflect.ValueOf(result)
						sectionValue := resultValue.Elem().FieldByName(listconfname)
						sectionType := sectionValue.Type()
						fmt.Println(resultValue, sectionValue, listconfname, sectionType)
						//if reflect.ValueOf(result).FieldByName(key).Kind() == reflect.Struct {
						//} else {
						//	err = fmt.Errorf(fmt.Sprintf("The type of variable you pass in is %v and the type you need is struct!", reflect.ValueOf(result).Elem().FieldByName(key).Kind()))
						//}
					}
				} else {
					Color_fmt := color.New(color.FgYellow).Add(color.Underline)
					Color_fmt.Print("[WARN]: ")
					fmt.Printf("Skip line %d of the file and read the error\n", index+1)
					continue

				}

			}
		} else {
			err = fmt.Errorf(fmt.Sprintf("The type of variable you pass in is %v and the type you need is struct!", typeInfo.Elem().Kind()))
			//err = errors.New("The type of variable you pass in is %v and the type you need is struct!")
			return
		}
	} else {
		err = fmt.Errorf("Please pass in the address of the variable!")
		//err = errors.New("Please pass in the address of the variable!")
		return
	}
	return
}
