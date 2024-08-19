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

	fmt.Printf("The Current time is in Nano-Seconds %v\n", time.Now().Nanosecond())
	current_page_Size := os.Getpagesize()

	fmt.Printf("The Current Page Size is %v\n", current_page_Size)

	fmt.Printf("The Program Exited time is in Nano-Seconds %v\n", time.Now().Nanosecond())

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
