package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Printf("Respon Body: %v\n", res.Body)

	//eksekusi ioutil
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	sb := string(body)
	fmt.Println(sb)
}
