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
