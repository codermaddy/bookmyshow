package main

import (
  "fmt"
  "net/http"
  "os"
  "path/filepath"
  "html/template"
)

var tpl *template.Template

func index(w http.ResponseWriter, r *http.Request){
  tpl = template.Must(template.ParseGlob("template/*"))
  tpl.ExecuteTemplate(w, "index.html", nil)
}


func main(){
  mux := http.NewServeMux()

  workDir, _ := os.Getwd()
  staticPath := http.Dir(filepath.Join(workDir, "static"))

  files := http.FileServer(staticPath)

  mux.Handle("/static/", http.StripPrefix("/static",  files))
  mux.HandleFunc("/", index)

  http.ListenAndServe(":8080", mux)
}
