package main

import (
  "fmt"
  "net/http"
  "os"
  "path/filepath"
  "html/template"

  "example.com/authentication/authent"
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseGlob("template/*"))
}

func index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Welcome to BookMyShow")
}

func main(){
  mux := http.NewServeMux()

  workDir, _ := os.Getwd()
  staticPath := http.Dir(filepath.Join(workDir, "static"))

  files := http.FileServer(staticPath)

  mux.Handle("/static/", http.StripPrefix("/static",  files))
  mux.HandleFunc("/", index)
  mux.HandleFunc("/login", authent.Login)

  http.ListenAndServe(":8080", mux)
}
