package models

import(
  "fmt"
  "database/sql"
)

var SelectOneUser string = "SELECT * FROM user where email=? AND password=?"

type User struct {
  Name string
  Email string
  Passwd string
}

func UserFromSqlRow(row *sql.Row) *User{
  var user User
  err := row.Scan(&user.Name, &user.Email, &user.Passwd)
  if err == nil{
    return &user
  } else{
    fmt.Println(err)
    return nil
  }
}
