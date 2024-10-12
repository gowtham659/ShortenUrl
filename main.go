package main

import (
	"ShortenUrlApp/database"
	"ShortenUrlApp/handlers"
	"fmt"
	"log"
	"net/http"

	"html/template"

	"github.com/gorilla/mux"
)

func GetStatic(w http.ResponseWriter, r *http.Request){
    t,err:=template.ParseFiles("./static/index.html")
    if err!=nil{
        fmt.Println(err)
    }
    t.Execute(w,nil)
}

func main() {
    // Initialize the Oracle DB connection
    database.InitDB()

    // Initialize router
    r := mux.NewRouter()
    
    // API route for shortening URLs
    r.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")

    // Redirect short URL
    r.HandleFunc("/{shortCode}", handlers.RedirectURL).Methods("GET")
    // Serve static files (HTML, CSS, JS)
    r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

    // Start server
    log.Println("Server starting on port 8080...")
    log.Println("click here : http://localhost:8080/")
    log.Fatal(http.ListenAndServe(":8080", r))
}