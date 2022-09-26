# Summary Sesi 4 - Interface, Reflect, Go-routines & Waitgroup

## Interface
Interface adalah sebuah tipe data pada bahasa Go yang merupakan kumpulan dari definisi satu atau lebih method. Kita dapat menggunakan tipe data interface jika kita telah mengimplementasikan method-method yang telah didefinisikan oleh interface tersebut melalui tipe data lainnya.
<br>
Contoh:
```
type shape interfac {
  area() float64
  perimeter() float64
}
```
Untuk mencari tahu tipe data bisa gunakan:

> fmt.Printf("Type of c1: %T", c1")

### Type assertion
Ketika struct circle menambahkan satu method selain dari method-method yang didefinisikan oleh interface shape, maka variable c1 tidak dapat menggunakan method baru tersebut. Ini berfungsi untuk mengembalikan suatu tipe data interface kepada tipe data aslinya
<br>
Teknik type assertion:

> c1.(circle).volume()

Kita juga dapat memeriksa apakah sebuah type assertion yang kita gunakan berhasil atau tidak. Caranya dengan memberikan 2 variable opsional yang menjadi penampung dari hasil yang akan dikembalikan oleh type assertionnya.

> value, ok := c1.(circle)
