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
Pada looping bersarang, keyword break dan continue akan berlaku pada blok perulangan dimana ia digunakan saja.Adacara agar kedua keyword ini bisa tertuju pada perulangan terluar atau perulangan tertentu, yaitu dengan memanfaatkanteknik pemberian label.

