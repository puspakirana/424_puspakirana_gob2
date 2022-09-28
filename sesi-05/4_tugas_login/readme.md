### Install Package gopass

go mod init tugas
go get github.com/howeyc/gopass

### Validasi Akun
Validasi akun menggunakan operator OR (||) karena bisa jadi hanya salah satu dari data salah (username/password). Jika menggunakan operator AND (&&), tidak akan bisa mendeteksi kondisi tersebut (selalu False).

> if username != "Puspa" || p != "1234"