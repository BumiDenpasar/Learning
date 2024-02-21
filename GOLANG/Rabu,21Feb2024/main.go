package main

import "fmt"

// one return value function
func addNumber(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	return a + b
}

// two determined return value function
func persegiPanjang(p, l int) (keliling int, luas int) {
	keliling = (p * 2) + (l * 2)
	luas = p * l
	return
}

//unlimited parameters value (variadic function)
func _(numbers ...int) {
	fmt.Println("This is function with unlimited parameters value")
}

func main() {
	a := addNumber(0, 0)
	b := addNumber(2, 5)

	fmt.Println(a, b)

	// get two return value
	k, l := persegiPanjang(10, 10)
	fmt.Println(k, l)

	// jika tipe data dipanggil dengan () artinya itu untuk casting / mengubah tipe dari sebuah data
	aNumber := 5
	changeToFloat := float64(aNumber)
	fmt.Println(changeToFloat)
}
