package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func Snum(ass int64) (acc int64) {
	Sresult, _ := rand.Int(rand.Reader, big.NewInt(ass))
	return Sresult
}

func main() {
	for {
		abc := Snum(3)
		fmt.Println(abc)
		time.Sleep(time.Second)
	}
}
