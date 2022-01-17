package event

import (
  "net/http"
  "fmt"

  "example.com/authentication/authent"
  "example.com/models"
)

func BookTicket(w http.ResponseWriter, r *http.Request){
  ok, session := authent.LoggedIn(r); if ok{
    val, ok := r.URL.Query()["id"]; if ok && len(val) > 0{
      stmt, err := db.Prepare("SELECT * FROM events WHERE id=?")
      checkError(err)

      row := stmt.QueryRow(val[0])

      event := models.EventFromSqlQuery(row)
      if event == nil{
          http.Redirect(w, r, "/events", 302)
      }

      if r.Method == "GET"{
        tpl.ExecuteTemplate(w, "bookticket.html", event)
      } else{
        r.ParseForm()
        var errorMsg string
        
        if len(r.Form["attendee"]) == 0{
          errorMsg = "Attendee Name can't be empty"
        } else{
          stmt, err := db.Prepare("INSERT INTO booking(event_id, email_id, name, ticket_type) VALUES(?, ?, ?, ?)")
          checkError(err)

          _, err = stmt.Exec(val[0], session.Email, r.Form["attendee"][0], r.Form["ticketType"][0])
          checkError(err)
        }

        if errorMsg == ""{
          fmt.Println("Booked Ticket Succesfully")
          http.Redirect(w, r, "/tickets", 302)
        } else{
          tpl.ExecuteTemplate(w, "BookTicket.html", event)
        }
      }
    } else{
      http.Redirect(w, r, "/event", 302)
    }
  } else{
    http.Redirect(w, r, "/login", 302)
  }
}
