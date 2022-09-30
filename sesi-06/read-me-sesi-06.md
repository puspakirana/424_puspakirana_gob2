# Summary Sesi 6 - GO Micro, Web Services & Web Server Go Programming

## Web Server
### Web Server - Penjelasan code 1_web.go

Line 8 : variabel global bernama port untuk menyimpan nilai port localhost yang mengarah pada localhost:8080

Line 11: function HandleFunc berasal dari package http yang digunakan untuk keperluan routing. Function memiliki 2 parameter:
1. Yang pertama, untuk mendefinisikan path routingnya
2. Yang kedua, untuk menerima sebuah function dengan 2 parameter yaitu http.ResponseWriter dan pointer http.Request

> http.ResponseWriter adalah sebuah interface dengan berbagaimethod yang digunakan untuk mengirim response balikkepada client yang mengirimkan request.

Line 13: function ListenAndServe untuk menjalankan server aplikasi. Menerima 2 parameter, yaitu:
1. Keterangan port
2. http.Handler yang merupakan sebuah interface. Namun karena kita tidak menggunakan http.Handler,maka kita cukup memberikan zero value dari tipe data interfaceyaitu nil.

Line 16-19: function greet untuk mengirim response berupa tulisan "Hello World" kepada client.

Cara menjalankan aplikasi:
1. Tulis perintah go run 1_web.go pada terminal
2. Buka browser dan arahkan ke <http://localhost:8080/>

### API - Penjelasan code 2_api.go



## Gin Framework
