package main

import(
	"html/template"
	"fmt"
	"log"
	"net/http"
	"time"
)

type PageVariables struct{
	Date string
	Time string
}

func main(){
	http.HandleFunc("/", HomePage) /* handling sent request made to / */
	log.Fatal(http.ListenAndServe(":8000",nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now() // find the time right now
	HomePageVars := PageVariables{
		Date: now.Format("10-10-1010"),
		Time: now.Format("10:10:10"),
	}
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")

	t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
	    if err != nil { // if there is an error
	  	  log.Print("template parsing error: ", err) // log it
	  	}
	    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	    if err != nil { // if there is an error
	  	  log.Print("template executing error: ", err) //log it
	  	}
}
