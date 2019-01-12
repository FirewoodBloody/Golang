package main

import (
	"fmt"
	"strconv"
)

type ranges struct {
}

func ranges(a int64) ([]int, error) {

	s := strconv.FormatInt(a, 10)
	var sum []int64 = make([]int64, len(s))
	fmt.Printf("%T", s)
	for k, v := range s {
		num, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return nil, err
		}
		sum[k] = num
	}
}
