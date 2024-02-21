package main

import "fmt"

func main() {
	// deklarasi map dengan key tipe string, value tipe int
	var test = map[string]int{
		"test1": 1,
		"test2": 2,
	}
	fmt.Println(test["test1"]) // panggil map pake key tipe string

	for k, v := range test {
		fmt.Println(k, v)
	}

	delete(test, "test1") // delete(map name, key)
	fmt.Println(test)

	var _ = map[string]int{}             // without initial value
	var _ = make(map[string]int)         // deklarasi map dengan make
	var pointerMap = new(map[string]int) // deklarasi map pointer without initial value
	var _ = *pointerMap                  // access data dari sebuah pointer

}
