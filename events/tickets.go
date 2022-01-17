package event

import (
  "net/http"
  "time"

  "example.com/models"
  "example.com/authentication/authent"
)

func Tickets(w http.ResponseWriter, r *http.Request){
  ok, session := authent.LoggedIn(r); if ok{
    stmt, err := db.Prepare("SELECT * FROM events E, booking B WHERE E.id = B.event_id AND B.email_id=? AND E.time > ? ORDER BY E.time")
    checkError(err)

    rows, err := stmt.Query(session.Email, time.Now())
    checkError(err)
    defer rows.Close()

    var eventTickets [] map[string]interface{}

    for rows.Next(){
      event,ticket := models.EventAndTicketFromSqlQuery(rows)

      eventTickets = append(eventTickets, map[string]interface{}{"event": event, "ticket": ticket})
    }

    tpl.ExecuteTemplate(w, "tickets.html", eventTickets)
  } else{
    http.Redirect(w, r, "/login", 302)
  }
}
