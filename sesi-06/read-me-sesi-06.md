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

```
Cara menjalankan aplikasi:
1. Tulis perintah go run 1_web.go pada terminal
2. Buka browser dan arahkan ke http://localhost:8080/
```

### API - Penjelasan code 2_api.go

Line 36: method Header dari interfacehttp.ResponseWriter yang kemudian di chaining denganmethod Set. Dilakukan untuk menentukan bentuk dari dataresponse yang ingin kita kirimkan kepada client.

Line 37: mengkonversi data employees menjadidata berbentuk JSON untuk dikirimkan kepada client denganmenggunakan method NewEncoder yang berasal dari packagejson, yang kemudian kita chaining dengan method Encodeuntuk mengkonversi datanya menjadi bentuk JSON.

Line 41: Kemudian jika method yang dikirimkan oleh client bukanmethod GET, maka kita akan mengirimkan response errordengan menggunakan function Error dari package http.

Line 64-66: mendapatkan nilai inputform dari client dengan menggunakan method FormValueyang berasal dari struct http.Request.

Line 68: mengkonversi input age dari tipedata string menjadi tipe data int

Line 29: meregistrasikan functioncreateEmployee kedalam routingan

```
Testing di Postman:
1. Method: GET
URL: http://localhost:8080/employees
2. Method: POST
URL: http://localhost:8080/employee
```

**Mengirimkan data employee dengan dengantemplate html**
1. Membuat file template.html
2. Mengubah function getEmployees
Line 47: function template.ParseFilesyang berasal dari package html/template yang digunakanuntuk mem-parsing file html

Line 54: method bernama Executeyang digunakan untuk memberikan response kepada client berupa template html. Parameter kedua dari method Executedigunakan untuk mengirimkan data yang ingin kita tampilkanpada file html.

**Melooping data employee pada file template.html**


Format pernulisan range loop nya adalah {{range$namaVariable := .}}, yang dimana cara penulisan namavariablenya masih sama seperti biasa namun harus diawalidengan tanda dolar $.Lalu setelah itu kita juga perlu mengakhiri looping nya dengankeyword end yang juga perlu dimasukkan kedalam 2 tandakurung {{end}}.

```
Testing pada browser:

buka URL: http://localhost:8080/employees
```


## Gin Framework
Gin adalah sebuah framework untuk bahasa Go yang digunakan untuk keperluan http routing.

**Membuat project dengan gin framework**

1. Buat folder belajar-gin
> mkdir belajar-gin

2. Masuk ke folder tersebut
> cd belajar-gin

3. Buat go.mod
> go mod init belajar-gin

4. Instalasi gin
> go get -u github.com/gin-gonic/gin

5. Buat folder controllers dan routes

6. Buat file carController.go di dalam folder controllers

7. Buat file carRouter.go di dalam folder routers
