package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"unicode"

	"github.com/gorilla/mux"
)

type Person struct {
	Name   string
	Fruits []string
}

var PORT = ":8080"

var person = Person{
	Name:   "Puspa",
	Fruits: []string{"apel", "mangga", "pisang"},
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/puspa", greet)
	r.HandleFunc("/puspa/{msg}", greetMessage)
	log.Fatal(http.ListenAndServe(PORT, r))

}

func greet(w http.ResponseWriter, r *http.Request) {
	fruits := ""

	for _, v := range person.Fruits {
		fruits += v + " "
	}

	msg := "Nama:" + person.Name + " Fruits: " + fruits
	fmt.Fprint(w, msg)
}

func greetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	status := false
	for _, v := range params["msg"] {
		if unicode.IsSpace(v) {
			status = true
			break
		}
	}
	if status == false {
		check := reflect.TypeOf(params["msg"]).Kind()
		if _, err := strconv.Atoi(params["msg"]); err == nil {
			num, _ := strconv.Atoi(params["msg"])
			check = reflect.TypeOf(num).Kind()

		}

		if check == reflect.String {
			fmt.Fprint(w, "Hello saya dengan tipe data string")
		} else if check == reflect.Int {
			fmt.Fprint(w, "Hello saya dengan tipe data int")
		}
	} else {
		fmt.Fprint(w, "data yang ditampilkan mengandung spasi")
	}

}
