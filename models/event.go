package models

import (
  "fmt"
  "database/sql"
  "time"
)

type Event struct{
  Id int
  Name string
  Time time.Time
}

func EventFromSqlQuery(rows *sql.Rows)(*Event){
  var event Event
  //var tmpTime []uint8
  err := rows.Scan(&event.Id, &event.Name, &event.Time); if err == nil{
    //event.Time = tmpTime.Time;
    return &event
  } else{
    fmt.Println(err)
    return nil
  }
}
