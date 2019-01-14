package ranges

import (
	"fmt"
	"reflect"
	"strconv"
)

func DivisionInt(a interface{}) (Transformation []int, err error) {
	IntType := reflect.TypeOf(a)
	switch IntType.Kind() {
	case reflect.Int:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]int, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int(a)
		}

		return Transformation, nil
	case reflect.Int8:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]int, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int(a)
		}

		return Transformation, nil
	case reflect.Int16:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]int, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int(a)
		}

		return Transformation, nil
	case reflect.Int32:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]int, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int(a)
		}

		return Transformation, nil
	case reflect.Int64:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]int, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int(a)
		}

		return Transformation, nil
	case reflect.Uint:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]int, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int(a)
		}

		return Transformation, nil
	case reflect.Uint8:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]int, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int(a)
		}

		return Transformation, nil
	case reflect.Uint16:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]int, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int(a)
		}

		return Transformation, nil
	case reflect.Uint32:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]int, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int(a)
		}

		return Transformation, nil
	case reflect.Uint64:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]int, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int(a)
		}

		return Transformation, nil
	}
	return nil, nil
}
