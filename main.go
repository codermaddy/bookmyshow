package main

import (
  "net/http"
  "os"
  "path/filepath"
  "html/template"

  "example.com/authentication/authent"
  "example.com/event"
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseGlob("template/*"))
}

func index(w http.ResponseWriter, r *http.Request){
  http.Redirect(w, r, "/events", 302)
}

func main(){
  mux := http.NewServeMux()

  workDir, _ := os.Getwd()
  staticPath := http.Dir(filepath.Join(workDir, "static"))

  files := http.FileServer(staticPath)

  mux.Handle("/static/", http.StripPrefix("/static",  files))
  mux.HandleFunc("/", index)
  mux.HandleFunc("/events", event.Events)
  mux.HandleFunc("/book", event.BookTicket)
  mux.HandleFunc("/tickets", event.Tickets)
  mux.HandleFunc("/login", authent.Login)
  mux.HandleFunc("/logout", authent.Logout)

  http.ListenAndServe(":8080", mux)
}
