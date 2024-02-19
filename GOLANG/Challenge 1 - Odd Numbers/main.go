package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7}
	var oddNumber []int

	for _, number := range numbers {
		if number%2 != 0 {
			oddNumber = append(oddNumber, number)
		}
	}
	fmt.Println(oddNumber)
}
