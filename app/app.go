package app

import (
	"encoding/xml"

	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)
type user struct {
    UID   uint   `xml:"uid"`
    FName string `xml:"fname"`
    LName string `xml:"lname"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello serve")
}

func requestCustomer(writer http.ResponseWriter, r *http.Request) {
	u := []user{
		{6089, "thewan", "perera"},
	}
	writer.Header().Add("Content-Type", "application/xml")
	xml.NewEncoder(writer).Encode(u)

}
func getCustomer(w http.ResponseWriter,r*http.Request){
	v:=mux.Vars(r)
	fmt.Fprintf(w,v["customer_id"])

}

func Start(){
	r:=mux.NewRouter()
	r.HandleFunc("/hello", hello).Methods(http.MethodGet)
	r.HandleFunc("/customer", requestCustomer)
	r.HandleFunc("/customer/{customer_id:[0-9]+}",getCustomer)

	// server starting point
	http.ListenAndServe(":8080", r)
}