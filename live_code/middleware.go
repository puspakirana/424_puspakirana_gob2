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

var person = Person{
	Name:   "Puspa",
	Fruits: []string{"apel", "mangga", "pisang"},
}

func main() {
	mux := http.NewServeMux()
	endpoint := http.HandlerFunc(greet)
	mux.Handle("/", middleware1(middleware2(endpoint)))

	fmt.Println("Listening to port 3000")
	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	status := false
	for _, v := range params["msg"] {
		if unicode.IsSpace(v) {
			status = true
			break
		}
	}
	if status == true {
		w.Write([]byte("data yang ditampilkan mengandung spasi"))
	} else {
		w.Write([]byte("Hello World!!!"))
	}

}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fruits := ""

		for _, v := range person.Fruits {
			fruits += v + " "
		}

		msg := "Nama:" + person.Name + " Fruits: " + fruits
		fmt.Println(msg)

		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		params := mux.Vars(r)

		check := reflect.TypeOf(params["msg"]).Kind()
		if _, err := strconv.Atoi(params["msg"]); err == nil {
			num, _ := strconv.Atoi(params["msg"])
			check = reflect.TypeOf(num).Kind()

		}

		if check == reflect.String {
			fmt.Println("Hello saya dengan tipe data string")
		} else if check == reflect.Int {
			fmt.Println("Hello saya dengan tipe data int")
		}

		next.ServeHTTP(w, r)
	})
}
