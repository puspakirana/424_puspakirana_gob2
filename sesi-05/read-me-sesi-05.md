# Summary Sesi 5 - Channels, Defer & Exit , Error Handling

## Channels
Channel merupakan sebuah mekanisme untuk Goroutine saling berkomunikasi dengan Goroutine lainnya. Maksud berkomunikasi disini adalah proses pengiriman dan pertukaran data antara Goroutine satu dengan Goroutine lainnya.

### Channel Operators
Kita membutuhkan operator dari channel agar channel tersebut dapat dijadikan sebagai alat untuk berkomunikasi antara Goroutine dengan yang lainnya.
<br>
Contoh:
```
func main() {
  c := make(chan string)
  
  //Mengirim data kepada channel
  c <- value
  
  //Menerima data dari channel
  result := <- c
}
```
Tanda <- merupakan sebuah operator dari channel untuk proses pengiriman data dari Goroutine satu dengan yang lainnya. 

### Function Close pada implementasi channel untuk proses komunikasi goroutine
Function close merupakan sebuah function yang digunakan untuk mengindikasikan bahwa sebuah channel sudah tidak digunakan untukberkomunikasi lagi. Perlu diingat bahwa cara ini perlu dilakukan.

### Channel directions
Ketika kita menggunakan channel sebagai sebuah parameter dari sebuah function, kita dapat menentukan apakah channel tersebut digunakan untuk menerima data saja, mengirim data saja, ataupun menerima sekaligus mengirim data. Channel directions ini memiliki sifat yang opsional, dan digunakan untuk kepentingan type-safety. Gambar dibawah ini merupakan penjelasan dari channel direction.
<br>
| Parameter Syntax  | Detail   |
|---|---|
| c chan string  | parameter c dapat digunakan untuk mengirim dan menerima data  |
| c chan <- string | parameter c hanya dapat digunakan untuk mengirim data  |
| c <- chan string  | parameter c hanya dapat digunakan untuk menerima data |

### Buffered vs unbuffered channel
Pada dasarnya, channel bersifat unbuffered atau tidak di buffer di memori. Channel yang kita gunakan pada halaman-halaman sebelumnya merupakan unbuffered channel yang dimana proses penerimaan dan pengiriman data bersifat synchronous. Lain halnya dengan unbuffered channel yang dimana kita dapat menentukkan kapasitas dari buffer nya, dan selama jumlah data yang dikirim melalui unbuffered channel tidak melebihi kapasitasnya, maka proses pengiriman data akan bersifat asynchronous.

### Select
Select merupakan sebuah fitur pada bahasa Go yang memungkinkan kita untuk dapat menggunakan lebih dari satu channel untuk proses komunikasi antara Goroutine satu dengan yang lainnya. 

## Defer & Exit

### Defer
Defer merupakan sebuah keyword pada bahasa Go yang digunakan untuk memanggil sebuah function yang dimana proses eksekusi akan di tahan hingga block sebuah function selesai.

### Exit
Funcion exit yang yang berasal dari package os berguna untuk menghentikan suatu program secara paksa. Perhatikan hasil eksekusi syntax dari gambar pertama pada gambar kedua.

## Error, Panic, & Recover

### Error
Error merupakan sebuah tipe data pada bahasa Go yang digunakan untuk me-return sebuah error jika ada error yang terjadi. Umumya, nilai error akan di-return pada pada posisi terakhir dari nilai-nilai return sebuah function.
<br>
- Custom Error
Kita juga dapat membuat pesan error kita sendiri dengan menggunakan method New yang merupakan sebuah method miliki tipe data error. 
<br>
Contoh:

> errors.New("Password has to have more than 4 characters"

## Panic
Panic digunakan untuk menampilkan stack traceerror sekaligus menghentikan flow goroutine(karena main() juga merupakan goroutine, makabehaviour yang sama juga berlaku). Setelah adapanic, proses akan terhenti, apapun setelah tidakdi-eksekusi kecuali proses yang sudah di-defersebelumnya (akan muncul sebelum panic error).Panic menampilkan pesan error di console, samaseperti fmt.Println(). Informasi error yangditampilkan adalah stack trace error, jadi sangatmendetail dan heboh.

## Recover
Function recover digunakan untuk menangkaperror yang terjadi pada sebuah Goroutine.Sekarang kita akan menambahkan sebuahfunction bernama catchErr  pada codingan kitasebelumnya yang akan menangkap error panicnya menggunakan function recover.Dan perlu  diingat bahwa kita perlu memanggilfunction catchErr dengan keyword defer.
