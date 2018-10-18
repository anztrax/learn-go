package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"testGoProject/stringutil"
)

type Todo struct{
	Title string
	Done bool
}

type TodoPageData struct{
	PageTitle string
	Todos []Todo
}

type User struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
}

//basic middleware
func logging(f http.HandlerFunc)  http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		log.Println(r.URL.Path)
		f(w,r)
	}
}

func main(){
	// mux router
	r := mux.NewRouter();

	//html template
	filePrefix, _ := filepath.Abs("./static/")
	tmpl := template.Must(template.ParseFiles(filePrefix + "/layout.html"))

	r.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request){
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(w,"%s %s is %d years old !", user.Firstname, user.Lastname, user.Age)
	})

	r.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request){
		peter := User{
			"John",
			"Doe",
			25,
		}

		json.NewEncoder(w).Encode(peter)
	})

	r.HandleFunc("/", logging(func(w http.ResponseWriter, r *http.Request){
		data := TodoPageData{
			PageTitle:"My Todo List",
			Todos: []Todo{
				{Title: "task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})).Methods("GET")


	r.HandleFunc("/books/{title}/page/{page}", logging(func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You're requested the book : %s on the page %s\n", title, page)
	})).Methods("GET")

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static", http.StripPrefix("/static", fs))

	http.ListenAndServe(":8010", r)
	fmt.Printf("%s", stringutil.Reverse("hello there"));
}