# Practice Builde REST API

## How to Run App
1. Create database godb at MySQLQorkbench / phpMyAdmin
2. Go to the practice directory
3. Run the program

> go run main.go

4. Test using postman

## Postman Result
<https://github.com/puspakirana/424_puspakirana_gob2/tree/master/sesi-08/practice/postman>

## Error found during development

1. Folder name
```
Do not use number in the beginning
From: 4_practice -> To: practice
```

2. Module name
```
Better use same name as the folder
From: main -> To: practice
```

3. Database Connection
```
First, must create the db (godb) before run the program

In my case, I have XAMPP and MySQLWorkbench.
And the MySQLWorkbench works for me.
port: 3306
Don't forget to put password in this code if you have db password:

> root:password@tcp(127.0.0.1:3306)

Unless:

> root:mysql@123@tcp(127.0.0.1:3306)
```

4. False code
I miss type some codes, next time be more careful and recheck the code
