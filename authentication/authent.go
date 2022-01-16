package authent

import (
  "database/sql"
  "fmt"
  "net/http"
  "html/template"
  "strings"

  _ "github.com/go-sql-driver/mysql"
  "example.com/models"
)

var tpl *template.Template
var db *sql.DB

func init(){
  tpl = template.Must(template.ParseGlob("template/*"))
  db, _ = sql.Open("mysql", "him:h!m@nshU7@tcp(127.0.0.1)/bookmyshow")
}

func checkError(err error){
  if err != nil{
    fmt.Println(err)
  }
}

func Login(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET"{
    ok, session := LoggedIn(r); if ok{
      // redirect
      fmt.Println("Logged in", session.UUID)
      return
    } else{
      tpl.ExecuteTemplate(w, "login.html", nil)
    }
  } else{
    r.ParseForm()
    var user *models.User
    var errorMsg string

    if len(r.Form["username"]) == 0 || len(r.Form["password"]) == 0 || len(strings.TrimSpace(r.Form["username"][0])) == 0 || len(strings.TrimSpace(r.Form["password"][0])) == 0{
      errorMsg = "Username and Password Field can't be empty"
    } else{
      stmt, err := db.Prepare(models.SelectOneUser)
      checkError(err)

      row := stmt.QueryRow(r.Form["username"][0], r.Form["password"][0])
      user = models.UserFromSqlRow(row)

      if user == nil{
        errorMsg = "Incorrect Credentials"
      }
    }

    if errorMsg == ""{
      CreateSession(w, user.Email)
      fmt.Println("Loggedin succesfully")
      //redirect
    } else{
      tpl.ExecuteTemplate(w, "login.html", errorMsg)
    }

  }
}
