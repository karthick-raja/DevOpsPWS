package main

import (
	"net/http"
	"os"
	"strings"
	"log"
	"html/template"
	"github.com/karthick-raja/DevOpsPWS/pkg/devopsutil"
)

func validateFileHandler(w http.ResponseWriter, r *http.Request) {
	devopsutil.Print(w, "=======================================================")
    err := devopsutil.ValidateFile(w, r.FormValue("fileName"))
	devopsutil.Print(w, "=======================================================")
    if err != nil {
		devopsutil.Print(w, err.Error())
	} else {
		devopsutil.Print(w, "File validation success")
	}
}

func osCommandHandler(w http.ResponseWriter, r *http.Request) {
    
	cmd := strings.Split(r.FormValue("cmd"), " ")
	
	err := devopsutil.OScommand(w, cmd)
	
    if err != nil {
		devopsutil.Print(w, err.Error())
	} else {
		devopsutil.Print(w, "Command executed successfully")
	}	
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) { 
	t, err := template.ParseFiles("./Templates/MainPage.html") 
	if err != nil { 
		devopsutil.Print(w, err.Error()) 
	} 
	err = t.Execute(w, "") 
	if err != nil { 
		devopsutil.Print(w, err.Error()) 
	} 
} 

func main() {
    http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/ValidateFile", validateFileHandler)
	http.HandleFunc("/OsCommand", osCommandHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}