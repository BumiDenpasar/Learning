package main

import "fmt"

func main() {
	/*
		numbers := "123"
		for index, value := range numbers {
			fmt.Println(index, value)
		}
	*/

	var names [4]string

	for i := 0; i < 4; i++ {
		names[i] = string(i)
		fmt.Println(string(i))
	}

}
