# Summary Sesi 2

## Condition
Kondisional pada pemrograman dapat digunakan untuk mengatur alur dari suatu program. Pada bahasa Go, terdapat 2 macam kondisional yang dapat kita pakai yaitu if else dan switch
<br>
###Conditions (Temporary Variable)
Variable tersebut hanya bisa diaksesdan digunakan pada scope block dari suatu kondisional.
<br>
Contoh:

> if **age:= currentYear - 1998**; age < 17 { ... }

### Switch
Switch juga digunakan untuk membuat suatu kondisional seperti if else. Switch pada bahasa Go tidak memerlukan keyword break,jadi jika suatu case  telah terpenuhi maka dia tidak akanberlanjut kepada case berikutnya.

### Switch with relational operators
Switch juga dapat digunakan dengan operator perbandingan.
<br>
Contoh:
> case (score < 8) && (score > 3):

### Switch (fallthrough keyword)
Falltrough digunakan ketika kita ingin switch melanjutkan pengecekan kepada case selanjutnya walaupun suatu case telah terpenuhi

###Nested Conditions
Kondisional bersarang merupakan sebuah proses kondisional yangdidalamnya terdapat proses kondisional kembali. Kita dapatmenggunakan if, else if , else, switch ataupun menggabungkanseluruhnya.

##Loopings
Looping atau perulangan merupakan suatu proses berulang yang dimana proses tersebut akan berhenti jika memenuhisuatu kondisi.

### First way of looping
> for i := 0; i < 3; i++ { ... }

### Second way of looping
Dengan menyelipkan kondisional seperti padalooping dengan menggunakan keyword while.
<br>
Contoh:
<br>
```
var i = 0
for i < 3 {
  ...
  i++
}
```

### Third way of looping
Dengan melakukan looping tanpa memberikan kondisi apapun.
<br>
Contoh:
<br>
```
var i = 0
for {
  ...
  i++
  if i == 3 {
    ...
   }
}
```

### Loopings (break and continue keywords)
Break -> untuk menghentikan sebuah proses looping
<br>
Continue -> untuk melanjutkan suatu proses looping

### Loopings (Nested Looping)
Nested looping atau looping bersarang adalah suatu proses looping yang memiliki suatu proses looping di dalamnya.

### Loopings (Label)
Pada looping bersarang, keyword break dan continue akan berlaku pada blok perulangan dimana ia digunakan saja.Ada cara agar kedua keyword ini bisa tertuju pada perulangan terluar atau perulangan tertentu, yaitu dengan memanfaatkanteknik pemberian label.

## Array
Array merupakan sebuah tipe data untukmenyimpan sebuah kumpulan dari data-data. Data-data yang kitasimpan pada sebuah array dalam bahasa Go harus memiliki tipedata yang sama, kecuali kita menyimpannya sebagai suatu interface kosong.
<br>
Array pada bahasa Go memiliki sifat fixed-length atau memilikipanjang yang tetap yang harus kita tentukan dari awal kitamembuat arraynya.
<br>
Kita dapat membuat array dengan 2 macam cara yaitu:
1. Dengan mendeklarasikannya terlebih dahulu tanpa memberi nilai apapun

```
var numbers [4]int
numbers = [4]int{1, 2, 3, 4}
```

2. Mendeklarasikan dan langsung menginisialisasikannya dengan memberikannya sebuah nilai.

```
var numbers = [3]int{1, 2, 3}
```

### Modify Element Through Index
Kita juga dapat memodifikasi element-element yang terdapatdalam sebuah array dengan cara mengakses indexnya.
<br>
Contoh:
<br>
> fruits[0] = "mango"

### Loop through elements
Ada 2 cara penulisan agar kita dapat melakukan looping untukmengakses element-element yang terdapat pada sebuah array.Perhatikan pada gambar pertama disebelah kanan. 
1. Menggunakan range loop

> for i, v in range fruits { ... }

2. Menggunakan looping biasa

> for i:= 0; i < len(fruits); i++ { ... }

### Multidimensional Array
Contoh:
<br>
> balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

## Slice
Slice merupakan suatu tipe data yang mirip dengan tipe data array, yang juga memiliki kegunaan untuk menyimpan satu ataulebih data. Namun tipe data slice dan array memiliki sifat yang berbeda. Slice tidak memiliki sifat fixed-length  yang berartipanjang dari slice tidak tetap sehingga kita bisa dengan leluasa menentukan panjang dari slice nya.

### Make function
Function make digunakan untuk membuat slice.
<br>
Contoh:
> var fruits = make([]string, 3)

### Append function
Jika kita ingin menambahkan element pada slice dari variablefruits pada halaman sebelumnya, kita bisa melakukannyadengan cara mengakses indexnya. Namun, jika ingin lebih mudah, maka kita bisa menggunakan fungsi **append**
<br>
Contoh:
<br>
> fruits = append(fruits, "apple", "banana", "mango")

### Append function with ellipsis
Jika kita ingin memasukkan seluruh element-element padasuatu array ke dalam array lainnya, maka kita dapatmenggunakan tanda ellipsis (...) atau tanda titik tiga berurut.
<br>
Contoh:
<br>
> fruits1 = append(fruits1, fruits2...)

### Copy function
Untuk mengcopy seluruh element pada sebuah slice ke dalam slice lain.
<br>
Contoh:
> nn := copy(fruits1, fruits2)

### Slicing
Digunakan untuk mendapatkan element-element dari sebuah slice dan bisa menentukan element dari index ke berapa.
<br>
Contoh:
<br>
```
- Mendapatkan element index ke 1, 2, 3
var fruits2 = fruits1[1 : 4]
- Mendapatkan element mulai index 0 sampai index terakhir
var fruits3 = fruits1[0:]
- Mendapatkan element mulai index pertama (index 0) sampai index ke 2
var fruits4 = fruits1[:3]
- Mendapatkan seluruh element
var fruits5 = fruits1[:]
atau
var fruits5 = fruits1[:len(fruits1)]
```

### Combining slicing and append
Contoh:
<br>
> fruits1 = append(fruits1[:3], "rambutan")

### Backing Array
Setiap kita membuat suatu slice pada bahasa Go , secara otomatis Go akan membuat suatu array tersembunyi yang disebut dengan Backing array. Backing array akan bertugas untuk menyimpan element pada slice, bukan slice nya sendiri. Bahasa Go mengimplementasikan slice sebagai sebuah struktur data yang disebut dengan slice header. Slice headerterdiri dari:
- Alamat memori/address dari backing array.
- Panjang dari slice yang bisa didapatkan dari fungsi len.
- Kapasitas dari slice yang bisa didapatkan dari fungsi cap.

Ketika kita mencoba untuk mendapatkan  beberapa element dari sebuah slice yang sudah ada dengan cara melakukan slicing, maka Go tidak akan membuat suatu backing array baru melainkan slice tersebut akan berbagi backing array yang sama dengan slice yang sudah ada.

### Cap function
Fungsi cap dapat kita gunakan untuk mengetahui kapasitasdari sebuah array maupun slice.
<br>
Ketika kita melakukan slicing yang dimulai dari index ke-0 hingga index ke-y, maka kita akan mendapatkan element dari index ke-0 hingga element sebelum index ke-y dengan panjang yang sesuai dengan element yang kita dapatkan dan kapasitas yang sama dengan kapasitas slice aslinya. 
<br>
Sedangkan slicing yang dimulai dari index ke-x, yang dimana nilai x adalah lebih dari 0, membuat elemen ke-x slice yangdiambil menjadi elemen ke-0 slice baru.

### Creating a new backing array
Ketika kita ingin mendapatkan element-element dari slice yang sudah ada, namun kita juga ingin membuat backing array yang baru, maka kita dapat menggunakan fungsi append untuk melakukannya. 

## Map
Sama seperti tipe data array dan slice, map juga berfungsi untuk menyimpan satu atau lebih data namun, mapdisimpan sebagai key:value pairs (pasang key dan value).

### Looping with map
Ketika kita ingin melakukan loop terhadap sebuah map, maka kitadapat menggunakan range loop.
<br>
> for key, value := range person { ... }

### Deleting Value
Kita juga dapat menghapus value dari sebuah map dengan caramenggunakan fungsi delete.
<br>
> delete(person, "age")

### Detecting a value
Kita juga dapat mendeteksi keberadaan suatu value dengan caramengakses key dari map nya lalu memberikan 2 variable sebagai penampungnya.
```
value, exist := person["age"]

if exist {
  ... 
} else {
  ...
}
```

## Aliase
Aliase merupakan sebuah fitu pada bahasa Go yang digunakan sebagai nama alternative dari tipe data yangsudah ada. 
```
//Mendeklarasikan tipe data alias bernama second
 type second = uint
 var hour second = 3600
 fmt.Printf("Hour Type: %T\n", hours) // akan menghasilkan output uint
```
## String in Depth
Tipe data string pada Go terbentuk dari kumpulan tipe data byte yang di letakkan di dalam slice atau bisa kita sebutdengan slice of bytes. Tipe data byte pada Go merupakan tipe data alias dari tipe data uint8. Ketika kita melakukan indexing terhadap suatu string, maka kita akan mendapat nilai representasi dari byte nya.

### Looping Over String (byte-byte)
Looping ini akan menghasilkan angka - angka yang merupakan representasi dari angka desimal ASCII Code.
```
  city := "Jakarta"
	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}
```

### Looping Over String (rune-by-rune)
Ketika kita ingin mendapatkan panjang karakter dengan menggunakan function len(), maka sebetulnya kita tidaksedang mendapatkan panjang dari string berdasarkan karakternya, namun kita mendapatkan jumlah byte pada string.
```
  str1 := "makan"
	str2 := "mânca"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))
```



