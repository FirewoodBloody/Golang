package main

import "fmt"

var (
	coisn        = 50
	user         = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adirano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(user))
)

func main() {
	var sum int
	for _, value := range user {
		for _, value1 := range value {
			switch string(value1) {
			case "a", "A", "e", "E":
				distribution[value] += 1
			case "i", "I":
				distribution[value] += 2
			case "o", "O":
				distribution[value] += 3
			case "u", "U":
				distribution[value] += 5
			}
		}
	}
	for index, value := range distribution {
		fmt.Printf("%s : %d\n", index, value)
		sum += value
	}
	fmt.Printf("%d", coisn-sum)
}
