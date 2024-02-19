package main

import "fmt"

func main() {
	var name string = "bumi"
	var name2 = "bumi"
	// KALAU PAKE KOMA DIA OTOMATIS NAMBAHIN SPASI
	fmt.Println(name, name2)
	fmt.Println(name + name2)

	name = "dewa"

	// ini adalah variable yang tidak perlu tipe data
	tanpaTipe := "test"
	fmt.Println(tanpaTipe)
	fmt.Println(name)

	// variabel "reserved" biasanya dipakai sebagai penampungan data yang tidak dipakai
	_ = "testing"
	_ = "testing2"
	test, _ := "testingedit", "testingedit2"
	fmt.Println(test)

	point := new(string)

	fmt.Println("pointer:", point, "nilai variablenya:", *point)
}
