package iniconfig

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

var listconfname string

func MarshalFile(data interface{}, filename string) (err error) {
	resule, err := Marshal(data)
	if err != nil {
		return
	}
	ioutil.WriteFile(filename, resule, 0755)
	return nil
}

func UnMarshalFile(filename string, result interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err == nil {
		return UnMarshal(data, result)

	} else {
		return err
	}

	return nil
}

func Marshal(data interface{}) (result []byte, err error) {
	var datas []string
	valueInfo := reflect.ValueOf(data)
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() == reflect.Struct {
		for i := 0; i < typeInfo.NumField(); i++ {
			if typeInfo.Field(i).Type.Kind() == reflect.Struct {
				TagVal := typeInfo.Field(i).Tag.Get("ini")
				if len(typeInfo.Field(i).Tag.Get("ini")) == 0 {
					TagVal = typeInfo.Field(i).Name
				}
				datas = append(datas, fmt.Sprintf("\n[%v]\n", TagVal))
				//fmt.Println(datas)
				for j := 0; j < typeInfo.Field(i).Type.NumField(); j++ {
					FileTagVal := typeInfo.Field(i).Type.Field(j).Tag.Get("ini")
					if typeInfo.Field(i).Type.Kind() == reflect.Struct {
						if len(FileTagVal) == 0 {
							FileTagVal = typeInfo.Field(i).Type.Field(j).Name
						}
						datas = append(datas, fmt.Sprintf("%v=%v\n", FileTagVal, valueInfo.Field(i).Field(j).Interface()))
					} else {
						err = fmt.Errorf(fmt.Sprintf("The type of variable you pass in is %v and the type you need is struct!", typeInfo.Elem().Kind()))
						//err = errors.New("The type of variable you pass in is %v and the type you need is struct!")
						continue
					}
				}
			} else {
				err = fmt.Errorf(fmt.Sprintf("The type of variable you pass in is %v and the type you need is struct!", typeInfo.Elem().Kind()))
				//err = errors.New("The type of variable you pass in is %v and the type you need is struct!")
				continue
			}
		}
	} else {
		err = fmt.Errorf("Please pass in the address of the variable!")
		//err = errors.New("Please pass in the address of the variable!")
		return
	}
	for _, val := range datas {
		result = append(result, []byte(val)...)
	}

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
								listconfname = typeInfo.Elem().Field(i).Name
								//fmt.Println(listconfname)
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
						key := strings.TrimSpace(value[0:posiTion])
						val := strings.TrimSpace(value[posiTion+1:])
						resultValue := reflect.ValueOf(result)
						sectionValue := resultValue.Elem().FieldByName(listconfname)
						sectionType := sectionValue.Type()
						//fmt.Println(sectionType)
						if sectionType.Kind() == reflect.Struct {
							KeyfindName := ""
							for i := 0; i < sectionType.NumField(); i++ {
								if sectionType.Field(i).Tag.Get("ini") == key {
									KeyfindName = sectionType.Field(i).Name
									if len(KeyfindName) != 0 {
										fiedvalue := sectionValue.FieldByName(KeyfindName)
										if fiedvalue != reflect.ValueOf(nil) {
											switch fiedvalue.Type().Kind() {
											case reflect.String:
												fiedvalue.SetString(val)
												break
											case reflect.Int:
												ints, err := strconv.ParseInt(val, 10, 64)
												if err == nil {
													fiedvalue.SetInt(ints)
												} else {
													Color_fmt := color.New(color.FgYellow).Add(color.Underline)
													Color_fmt.Print("[WARN]: ")
													fmt.Println("String conversion int failed")
													continue
												}

											}
										} else {
											Color_fmt := color.New(color.FgYellow).Add(color.Underline)
											Color_fmt.Print("[WARN]: ")
											fmt.Printf("Skip line %d of the file and read the error\n", index+1)
											continue
										}
									} else {
										Color_fmt := color.New(color.FgYellow).Add(color.Underline)
										Color_fmt.Print("[WARN]: ")
										fmt.Printf("Skip line %d of the file and read the error\n", index+1)
										continue

									}

								}
							}
						} else {
							err = fmt.Errorf(fmt.Sprintf("The type of variable you pass in is %v and the type you need is struct!", sectionType.Kind()))
							return
						}
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
