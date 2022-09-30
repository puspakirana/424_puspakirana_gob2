### Install Package gopass

```
go mod init tugas
go get github.com/howeyc/gopass
```

### Account Validation
1. Checking with not equals to (!=) relational operator
When we use not equals to (!=), account validation must use OR (||) logical operator because there might be a case that user input only one wrong data (either username/password). If we use AND (&&) logical operator for that condition, it will always return False.

> if username != "Puspa" || p != "1234"

2. Checking with equals to (==) relational operator
When we use equals to (==), the operator that should be used is AND (&&) logical operator, because both data (username & password) must be true. AND (&&) logical opertaor will return False if one or both condition are not fulfilled.

> if username == "Puspa" && p == "1234"

Note:
> Logical operators are related to truth table. Learn more here: <https://en.wikipedia.org/wiki/Truth_table>

### How to run the program
1. Open terminal
2. Go to .go file directory (in this case 4_tugas_login folder)

> cd 4_tugas_login

3. Run the program

> go run main.go

4. Input the username, press Enter
<img width="151" alt="Screen Shot 2022-09-28 at 22 24 57" src="https://user-images.githubusercontent.com/76604032/192820250-bd2f4fe3-2b5a-4efd-a646-404665c06562.png">

5. Input the password, press Enter

<img width="178" alt="Screen Shot 2022-09-28 at 22 25 30" src="https://user-images.githubusercontent.com/76604032/192820371-e254640d-eb74-497f-8513-9f3d3b37579a.png">

6. And the result will come up:
- Account is validated

<img width="271" alt="Screen Shot 2022-09-28 at 22 25 08" src="https://user-images.githubusercontent.com/76604032/192820438-4aff38b8-5cb9-4bd7-9f82-62944ec87d39.png">

- Account is not validated

<img width="370" alt="Screen Shot 2022-09-28 at 22 25 39" src="https://user-images.githubusercontent.com/76604032/192820489-7898abb6-07cb-433c-969c-5f1f4bc3158b.png">
