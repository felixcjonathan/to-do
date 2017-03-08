package main

import (
	"html/template"
	"log"
	"net/http"
)

var toDoList []string

func main() {
	port := "8081"

	http.HandleFunc("/", index)
	http.HandleFunc("/add", add)
	log.Println("App started...")
	http.ListenAndServe(":"+port, nil)
}

// sayHello will directly write to response writer
func add(w http.ResponseWriter, r *http.Request) {
	toDo := r.FormValue("to_do")
	if toDo != "" {
		toDoList = append(toDoList, toDo)
	}
	//w.Write([]byte(toDo))
	http.Redirect(w, r, "/", http.StatusFound)
}

// index will parse template file and write it to response writer
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, toDoList)
}
