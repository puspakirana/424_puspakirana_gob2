# Summary Sesi 3 - Function, Method, Pointer, Struct & Exported-Unexported

##Function
Cara menulis function pada Go dengan menggunakan keyword *func* lalu diikuti dengan parameter yang digunakan
<br>
Contoh:
```
func greet(name string, age int) {
  ... }
```

### Function (Return)
Ketika function yang kita buat mengembalikan/me-return sebuahnilai, maka kita perlu menuliskan tipe data dari nilai yang di return.
<br>
Contoh:
```
func greet(name string, age int) string {
  ... 
  return ...
}
```

### Function (Returning multiple values)
Function pada bahasa Go dapat me-return lebih dari satu nilai.
<br>
Contoh:
<br> outputnya adalah area dan circumference
```
func calculate(d float64) (float64, float64) {
  ...
  return area, circumference
}
```

### Variadic function
Bahasa Go telah menyediakan sebuah function yang dimana functiontersebut dapat menerima argumen yang tak terbatas jumlahnya.
Function tersebut bernama function variadic. Dengan menambahkan tanda titik tiga pada parameter pada sebuahparameter, maka sebuah dianggap sebagai sebauh function variadic,dan parameter yang diterminya akan di konversikan menjadi sebuah slice.
<br>
Contoh:
```
func print(names ...string) []map[string]string {
  var result []map[string]string
  ...
  return result
}
```

## Closure
Closure merupakan merupakan sebuah anonymous function atau function tanpa nama yang dapat disimpan sebagai sebuahvariable maupun dapat dijadikan sebagai nilai return dari sebuah function.

### Closure (Declare closure in variable)
Contoh untuk membuat closure yang dapat disimpan sebagai variable:
```
var evenNumbers = func(numbers ...int) []int {
  ...
}
```

### Closure (IIFE)
IIFE atau singkatan dari immediately-invoked function expression merupakan sebuah closure yang dapat langsung tereksekusi ketika pertama kali dideklarasikan. Jika kita ingin membuat suatu closure menjadi IIFE, maka kita perlu langsung memanggil closure tersebut secara langsung pada saat dideklarasikan. Perlu diingat bahwa kita tidak perlu lagi memanggil closure IIFE dengan tanda kurung () karena closure IIFE tereksekusi pada saat dideklarasikan.
<br>
Contoh:
```
var isPalindrome = func(str string) bool {
  ...
} ("katak")
fmt.Println(isPalindrome)
```

### Closure as a return value
Closure juga bisa dijadikan sebagai nilai kembalian dari suatu function.

### Closure (Callback)
Callback adalah sebuah closure yang dijadikan sebagai parameter pada sebuah function. 

## Pointer
Pointer merupakan sebuah variable spesial yang digunakan untuk menyimpan alamat memori variable lainnya. Sebagaicontoh sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori dimana nilai 4disimpan, bukan nilai 4 itu sendiri.Variabel-variabel yang memiliki reference atau alamat memori yang sama, saling berhubungan satu sama lain dan nilainyapasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain (yang alamat memorinya sama) yaitu nilainya juga ikut berubah.
<br>
Contoh:
```
var first *int
var second *int
_, _ = first, second
```

### Memory Address
Untuk menampilkan alamat memori dari variable dilakukan dengan cara menggunakan tanda ampersand (&).
- Kita dapat mendapatkan alamat memori dari variable biasa dengan menggunakan tanda ampersand (&).
- Kita juga dapat mendapatkan nilai asli dari variable pointer dengan cara menggunakan tanda asterisk (*).

> fmt.Println(&second)

### Changing value through pointer
Pointer digunakan untuk menyimpan alamat memori, maka ketika kita mengganti nilai dari sebuah pointer, maka variable lainnya yang mempunyai alamat memori yang sama, akan ikut terganti nilainya.

### Pointer as parameter
Pointer juga bisa dijadikan sebagai sebuah parameter pada sebuah function.
```
//function
func changeValue(number *int){
  ...
}

//penggunaan function
changeValue(&a)
```

### Struct
Bahasa Go tidak menyediakan tipe data class untuk membuat sebuah struktur data seperti layaknya bahasa pemrograman yang menganut paradigma OOP. Tetapi Go menyediakan sebuah tipe data yang bernama Struct untuk membuat sebuah struktur.
<br>
Struct adalah sebuah tipe data berupa kumpulan/koleksi dari berbagai macam property/field dan juga method. Cara membuat struct pada Go cukup mudah. Contohnya seperti pada gambar di sebelah kanan.
<br>
Contoh:
```
type Employee struct {
  name      string
  age       int
  division  string
}
```

### Giving value to struct
```
var employee Employee
employee.name = "Orang"
employee.age = 23
employee.division = "Developer"
```

### Initializing struct
Kita juga dapat menginisialisasi sebuah struct sekaligus memberikan nilai-nilai nya.
```
var employee = Employee{name: "Ananda", age: 20, division: "Finance"}
```

### Pointer to a struct
Kita juga dapat menggunakan pointer pada sebuah struct.
```
var employee2 *Employee = &employee1
```

### Embedded struct
Struct juga dapat mengandung tipe data struct lainnya dengan menjadikannya sebuah field.
```
type Person struct {
  name    string
  age     int
}

type Employee struct {
  division    string
  person      Person
}
```

### Anonymous struct
Anonymous struct adalah sebuah struct yang tidak dideklerasikan di awal sebagai sebuah tipe data struct baru, melainkan langsung dideklerasikan bersamaan dengan pembuatan variable.
```
//Anonymous Struct tanpa pengisian field

	var employee6 = struct {
		person   Person
		division string
	}{}
  
 //Anonymous Struct dengan pengisian field
var employee7 = struct {
  person   Person
	division string
}{
	  person:   Person{name: "Moana", age: 17},
		division: "Ocean Guardian",
}
```

### Slice of struct
Slice juga dapat dikombinasikan dengan tipe data struct, cara penulisannya mirip seperti slice of map.
```
var people = []Person{
	{name: "Belle", age: 18},
	{name: "Jasmine", age: 15},
	{name: "Sofia", age: 10},
}
```

### Slice of anonymoius struct
Anonymous struct juga dapat digabungkan dengan tipe data slice dan pengisian nilainya pun dapat dilakukan secara langsung.
```
var employee8 = []struct {
	person   Person
	division string
}{
	{person: Person{name: "Belle", age: 18}, division: "Personal Development"},
	{person: Person{name: "Jasmine", age: 15}, division: "Magic Carpet Training"},
	{person: Person{name: "Sofia", age: 10}, division: "Friendship and Relation"},
}
```

## Exported & Unexported
Pada bahasa Go, tiap folder yang berbeda akan dianggap sebagai suatu package. Kita dapat menggunakan berbagai variable ataupun tipe data dari package lainnya asalkan variable atau tipe data lainnya tersebut telah ter-eksport dari package nya. Kemudian cara kita meng-eksport suatu variable atau suatu tipe data adalah dengan mengawali penulisan variable maupun tipe data lainnya dengan **huruf kapital atau upper case**.

Satu hal lagi yang perlu kita ketahui yaitu terdapat suatu function yang akan dieksekusi terlebih dahulu sebelum function main. Function tersebut bernama *init.
Jika kita ingin menjalankan seluruh file yang berada di root direktori kita, maka kita menggunakan perintah go run *.go. 
