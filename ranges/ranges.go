package ranges

import (
	"fmt"
	"reflect"
	"strconv"
)

func Ints(a interface{}) (Transformation []interface{}, err error) {
	IntType := reflect.TypeOf(a)
	switch IntType.Kind() {
	case reflect.Int:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]interface{}, len(s))
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
		Transformation := make([]interface{}, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int8(a)
		}

		return Transformation, nil
	case reflect.Int16:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]interface{}, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int16(a)
		}

		return Transformation, nil
	case reflect.Int32:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]interface{}, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int32(a)
		}

		return Transformation, nil
	case reflect.Int64:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]interface{}, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = int64(a)
		}

		return Transformation, nil
	case reflect.Uint:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]interface{}, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = uint(a)
		}

		return Transformation, nil
	case reflect.Uint8:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]interface{}, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = uint8(a)
		}

		return Transformation, nil
	case reflect.Uint16:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]interface{}, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = uint16(a)
		}

		return Transformation, nil
	case reflect.Uint32:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]interface{}, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = uint32(a)
		}

		return Transformation, nil
	case reflect.Uint64:
		s := fmt.Sprintf("%v", a)
		Transformation := make([]interface{}, len(s))
		for k, v := range s {
			a, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return nil, err
			}
			Transformation[k] = uint64(a)
		}

		return Transformation, nil
	}
	return nil, nil
}
