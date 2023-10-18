package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("THis is a ,mux gorillallal")
	greet()
r:=mux.NewRouter()
r.HandleFunc("/",serve).Methods("GET")

log.Fatal(http.ListenAndServe(":3000",r))


}

func greet(){
	fmt.Println("WElcome to mux ")
}
func serve(w http.ResponseWriter, r* http.Request){
	w.Write([] byte ("<h1> Welvome to the sever madeby golang</h1>"))
}