package models

import (
  "fmt"
  "time"
)

type Event struct{
  Id int
  Name string
  Time time.Time
}

type Scannable interface{
  Scan(dest ...interface{}) error
}

func EventFromSqlQuery(row Scannable)(*Event){
  var event Event
  err := row.Scan(&event.Id, &event.Name, &event.Time); if err == nil{
    return &event
  } else{
    fmt.Println(err)
    return nil
  }
}
