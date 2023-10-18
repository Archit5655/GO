package main

import (
	"fmt"
	"log"
	"net/http"

	"mongoapi.com/router"
)

func main() {
	fmt.Println("MOFGO")
r:=router.Router()
	fmt.Println("Server is strting")
log.Fatal(	http.ListenAndServe(":4000",r))
fmt.Println("Server running")
}
