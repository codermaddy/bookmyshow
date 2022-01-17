package models

import (
  "fmt"
)

type Ticket struct{
  BookingId int
  EventId int
  EmailId string
  TicketType string
  Attendee string
}

func TicketFromSqlQuery(row Scannable)(*Ticket){
  var ticket Ticket
  err := row.Scan(&ticket.BookingId, &ticket.EventId, &ticket.EmailId, &ticket.TicketType, &ticket.Attendee); if err == nil{
    return &ticket
  } else{
    fmt.Println(err)
    return nil
  }
}

func EventAndTicketFromSqlQuery(row Scannable)(*Event, *Ticket){
  var event Event
  var ticket Ticket
  err := row.Scan(&event.Id, &event.Name, &event.Time, &event.Organizer, &event.Venue, &ticket.BookingId, &ticket.EventId, &ticket.EmailId, &ticket.TicketType, &ticket.Attendee); if err == nil{
    return &event, &ticket
  } else{
    fmt.Println(err)
    return nil, nil
  }
}
