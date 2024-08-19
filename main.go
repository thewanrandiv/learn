package main

import (
	
	_ "encoding/json"
	"fmt"

	"net/http"

	"os"
	"time"
)

type user struct {
	uid   uint
	fname string
	lname string
}

func main() {
	// handle functions
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/customer", requestCustomer)

	// server starting point
	http.ListenAndServe(":8080", nil)



}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello serve")
}

func requestCustomer(w http.ResponseWriter, r *http.Request) {
	u := user{
		uid:   6089,
		fname: "Thewan",
		lname: "Perera",
	}

	fmt.Fprintf(w,"Response of Customer is" ,u)

}
