package authent

import (
  "net/http"

  _ "github.com/go-sql-driver/mysql"
  "example.com/models"
  "github.com/google/uuid"
)

func LoggedIn(r *http.Request) (bool, *models.Session){
  var session *models.Session
  uuidC, err := r.Cookie("user")
  uuid := uuidC.Value
  if err != nil{
    return false, nil
  } else{
    stmt, err := db.Prepare("SELECT * FROM session WHERE uuid=?")
    checkError(err)

    row := stmt.QueryRow(uuid)

    session = models.SessionFromSqlRow(row)
    if session != nil{
      return true, session
    } else{
      return false, nil
    }
  }
}

func CreateSession(w http.ResponseWriter, email string){
  uuid := uuid.NewString()
  http.SetCookie(w, &http.Cookie{Name: "user", Value: uuid, MaxAge: 0})


  stmt, err := db.Prepare("INSERT INTO session VALUES(?, ?)")
  checkError(err)

  _, err = stmt.Exec(uuid, email)
  checkError(err)
}
