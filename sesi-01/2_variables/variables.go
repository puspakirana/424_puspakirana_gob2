package main

import "fmt"

func main() {
	//belajar variable dengan data type
	var nama string
	var umur int = 21

	nama = "Puspa"
	umur = 22

	fmt.Println("Ini adalah namanya ==>  ", nama)
	fmt.Println("Ini adalah umurnya ==> ", umur)
	fmt.Println("-----------------")

	//belajar variable tanpa data type
	var nama2 = "Jane"
	var umur2 = 25

	fmt.Printf("%T,%T\n", nama2, umur2)
	fmt.Println("-----------------")

	//belajar short declaration
	namaSaya := "Ariel"
	umurSaya := 15

	fmt.Printf("%T,%T\n", namaSaya, umurSaya)
	fmt.Println("-----------------")

	//belajar multiple variable declarations
	var student1, student2, student3 string = "satu", "dua", "tiga"
	var first, second, third int
	first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)

	var name1, age1, address1 = "Aurora", 16, "Castle"
	data1, data2, data3 := 1, "Apple", 3.5

	fmt.Println(name1, age1, address1)
	fmt.Println(data1, data2, data3)
	fmt.Println("-----------------")

	//belajar underscore variable
	//untuk menghindari error jika suatu variable tidak digunakan
	var firstVariable string
	var name2, age2, address2 = "Jasmine", 18, "Somewhere over the rainbow"

	_, _, _, _ = firstVariable, name2, age2, address2
	fmt.Printf("test underscore variabel > %T\n", firstVariable)
	fmt.Printf("Halo nama saya %s, umur saya adalah %d tahun, saya tinggal di %s\n", name2, age2, address2)
	fmt.Println("-----------------")
}
