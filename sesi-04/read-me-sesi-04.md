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


## Empty Interface
Empty interface merupakan suatu tipe data yang dapat menerima segala tipe data pada bahasa Go. Maka dari itu, sebuah variable dengan tipe data empty interface dapat diberikan nilai dengan tipe data yang berbeda-beda.

## Reflect
Reflect digunakan untuk melakukan  inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkanmemanipulasinya. Dari banyak fungsi yang tersedia di dalam package tersebut, ada 2 fungsi yang paling penting untuk diketahui, yaitureflect.ValueOf() danreflect.TypeOf().
- Fungsi reflect.ValueOf() akan mengembalikan objek dalam tipe reflect.Value, yang berisikan informasi yangberhubungan dengan nilai pada variabel yang dicari
- Sedangkan reflect.TypeOf() mengembalikan objek dalam tipe reflect.Type. Objek tersebut berisikan informasi yangberhubungan dengan tipe data variabel yang dicari

## Concurrency - Goroutines
### Concurrency
Arti dari concurrency adalah mengeksekusi sebuah proses secara independen atau berhadapan dengan lebih dari satu tugas dalam waktu yang sama. Perlu diingat disini bahwa concurrency  berbeda dengan parallelism, karena parallelism memiliki arti mengerjakan tugas yang banyak secara bersamaan dari awal hingga akhir. Sedangkan pada concurrency, kita tidak akan tahu tentang siapa yang akan menyelesaikan tugasnya terlebih dahulu.
### Goroutines
Goroutine merupakan sebuah thread ringan pada bahasa Go untuk melakukan concurrency. Jika  dibandingkan dengan thread biasa, Goroutine memiliki ukuran thread yang jauh lebih ringan. Pada saat kita mengeksekusi sebuah Goroutine, maka satu Goroutine hanya membutuhkan 2kb memori saja, sedangkan satu thread biasa dapat menghabiskan 1-2mb memori.Goroutine bersifat asynchronous sehingga proses nya tidak saling tunggu dengan Goroutine lainnya. 


## Waitgroup
Waitgroup merupakan sebuah struct dari package sync, yang digunakan untuk melakukan sinkronisasi terhadap goroutine.

### Method Add
Method Add digunakan untuk menambah counter dari waitgroup.Counter dari waitgroup ini berguna untuk memberitahu waitgrouptentang jumlah go routine yang harus ditunggu.

### Method Done
Kemudian untuk memberitahu waitgroup tentang go routine yang telahmenyelesaikan proses nya, maka kita perlu memanggil method Done

### Method Wait
Method Waitberguna untuk menunggu seluruh go routine menyelesaikan proses nya,sehingga method Wait akan menahan function main hingga seluruhproses go routine selesai.

