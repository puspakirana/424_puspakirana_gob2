# Panduan

## User

1. UserRegister

Digunakan untuk menambah user

POST /user/register

### Response
<img width="295" alt="UserRegister" src="https://user-images.githubusercontent.com/76604032/196942752-15ad2253-ae22-4360-bc3a-fd4a14337597.png">

2. UserLogin

Digunakan user untuk mengakses akunnya

### Response
<img width="760" alt="UserLogin2" src="https://user-images.githubusercontent.com/76604032/196943401-1eb8a263-ada2-4f77-a1fd-ed6ba7be1056.png">


3. UpdateUser

- Digunakan user untuk update email, username, dan umur
- Disertai dengan autorisasi user (user harus login)

### Response
<img width="292" alt="UserUpdate" src="https://user-images.githubusercontent.com/76604032/196942764-075ab6eb-9c57-4bb1-af5f-6a79936bb7ac.png">

4. DeleteUser

- Digunakan user untuk menghapus akunnya
- Disertai dengan autorisasi user (user harus login)

### Response
<img width="465" alt="UserDelete" src="https://user-images.githubusercontent.com/76604032/196942779-0cf87865-8794-4ae4-8c60-db6ac212180e.png">

## Photo

1. CreatePhoto

- Digunakan user untuk memposting foto
- Disertai dengan autorisasi user (user harus login)

### Response
<img width="461" alt="CreatePhoto" src="https://user-images.githubusercontent.com/76604032/196944383-67fef8c3-ddf3-40e7-a897-15fb6ce5175b.png">

2. GetAllPhoto

- Digunakan user untuk mendapatkan semua data foto yang tersimpan di database aplikasi
- Disertai dengan autorisasi user (user harus login)

### Response
<img width="448" alt="GetAllPhoto" src="https://user-images.githubusercontent.com/76604032/196944389-101398ae-03ec-4177-b867-e6adde615e45.png">

3. UpdatePhoto

- Digunakan user untuk memperbarui foto
- Disertai dengan autorisasi user (user harus login)

### Response
<img width="443" alt="UpdatePhoto" src="https://user-images.githubusercontent.com/76604032/196944395-ea44e4e6-ec28-4120-92f7-bdad82af99ee.png">

4. DeletePhoto

- Digunakan user untuk menghapus foto
- Disertai dengan autorisasi user (user harus login)

### Response
<img width="488" alt="DeletePhoto" src="https://user-images.githubusercontent.com/76604032/196944398-07b1d240-6598-4259-84e4-8fd975514611.png">
