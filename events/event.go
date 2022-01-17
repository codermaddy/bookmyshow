package event

import (
  "fmt"
  "net/http"
  "database/sql"
  "html/template"
  "time"

  "example.com/authentication/authent"
  "example.com/models"
)

var tpl *template.Template
var db *sql.DB

func init(){
  tpl = template.Must(template.ParseGlob("template/*"))
  db, _ = sql.Open("mysql", "him:h!m@nshU7@tcp(127.0.0.1)/bookmyshow?parseTime=true")
}

func checkError(err error){
  if err != nil{
    fmt.Println(err)
    panic("Database Error")
  }
}

func Events(w http.ResponseWriter, r *http.Request){
  ok, _ := authent.LoggedIn(r); if ok{
    stmt, err := db.Prepare("SELECT * FROM events WHERE time > ?")
    checkError(err)

    rows, err := stmt.Query(time.Now())
    checkError(err)
    defer rows.Close()

    var events []*models.Event
    for rows.Next(){
      event := models.EventFromSqlQuery(rows)
      events = append(events, event)
    }

    tpl.ExecuteTemplate(w, "events.html", events)
  } else{
    http.Redirect(w, r, "/login", 302)
  }
}
