package models

import (
  "fmt"
  "database/sql"
)

type Session struct{
  UUID string
  Email string
}

func SessionFromSqlRow(row *sql.Row) (*Session){
  var session Session
  err := row.Scan(&session.UUID, &session.Email)
  if err == nil{
    return &session
  } else{
    fmt.Println(err)
    return nil
  }
}
