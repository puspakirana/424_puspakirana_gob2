

## Web Service
Web service adalah penghubung frontent dan backend atau bisa disebut RESTFUL API

## Go Installation

### Add path golang mac
---
>cd ~
><br>
>echo "export GOPATH=/Users/arif/Documents/go" >> .bash_profile
><br>
>echo "export GOPATH=/Users/arif/Documents/go" >> .zshrc
><br>
>cat .bash_profile
><br>
>cat .zshrc
><br>
>echo $GOPATH

### Cek versi golang
----
>go version
><br>
>echo $GOPATH

### Create folder project
---
>mkdir hello
><br>
>cd hello

### Instalasi setup
---

tools: vscode (visual studio code)
<br>
plugins/extensi:
- code runner
- golang / go
- Prettier
- Code formatter
tools golang: https://go.dev/dl/

## Introduction to Go
Creating First Project: Hello World

### First project
---
untuk create module
> go mod init hello

create file golang
> touch main.go

run golang
> go run main.go

untuk build exe membutuhkan module
> go build -o main.exe 

- Dalam sebuah proyek, harus ada function bernama main. Function in akan dipanggil pertama kali pada saat eksekusi program
- Function fmt.Println() digunakan untuk memunculkan text ke layar (terminal) dan menghasilkan baris baru di akhir outputnya
- Function fmt.Print() kegunaannya sama dengan fmt.Println(). Hanya saja tidak menghasilkan baris baru di akhir.
- Function fmt.Printf untuk mengetahui tipe data suatu variable
- Komentar multiline pada go diawali dengan /* dan diakhiri */
- Komentar inline pada go diawali dengan //

## Variables
- with data type
- without data type
- multiple declaration
- underscrore variable, kegunaannya untuk menghindari error yang akan terjadi jika ada suatu variable yang menganggur atau tidak dipakai

Function fmt.Printf bisa digunakan untuk memasukkan berbagai macam variable dengan tipe data berbeda-beda ke dalam satu kalimat

## Data Types

### 4 Kategori Tipe Data Golang
1. Basic Type: number, string boolean
2. Aggrefate Type: array & struct
3. Reference Type: slice, pointer, map, function, channel
4. Interface Type: interface

### Basic Type

**Integer:** tipe data numerik non desimal. Terbagi menjadi:
- int : bilangan cacah (positif) 
- uint : bilangan bulat (positif & negatif)

**Float:** tipe data numerik desimal. Dibagi jadi dua, float32 dan float64

**Bool:** tipe data yang digunakan dalam perkondisian atau perulangan. Bernilai binary (true/false, 1/0, etc)

**String:** tipe data yang nilainya diapit oleh tanda quote atau petik dus ("")

**Nil:** bukan tipe data, melainkan sebuah nilai kosong. Nil tidak bisa digunakan pada tipe data yang sudah dibahas sebelumnya. Tapi, bisa diset ke tipe data:
- pointer
- tipe data function
- slice
- map
- channel
- interface kosong atau interfac{}

**Zero:** berbeda dengan nil. Semua tipe data yang dibahas sebelumnya memiliki zero value (nilai default tipe data). Contoh:
- Zero value dari string adalah "" (string kosong).
- Zero value dari bool adalah false.
- Zero value dari tipe numerik non-desimal adalah 0.
- Zero value dari tipe numerik desimal adalah 0.0.
- Zero value berbeda dengan nil. Nil adalah nilai kosong, benar-benar kosong.

## Constant Variable
atau konstanta, merupakan jenis variable yang nilainya tidak dapat diubah. Contoh: PI, kecepatan cahaya, etc

## Operators

**Operator Aritmatika:** tambah, kurang, kali, bagi, sisa hasil bagi, increment, decrement

**Operator Perbandingan:** equals to, not equals to, greater than, less than, greater than equals to, less than equals to

**Operator Logika:** And, Or, Not
