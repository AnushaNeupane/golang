package main

import(
	"fmt"
	"log"
	"net/http"
)


func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a request on /")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main(){
	http.HandleFunc("/", handlerFunc)
	log.Fatal(http.ListenAndServe(":8000",nil))
}