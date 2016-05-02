package main

/**
 * Simple HTTP web server to practice golang
 * @juanvallejo
 */

import (
	"fmt"
	"net/http"
	"log"
)

func sayhelloWorld(writer http.ResponseWriter, request *http.Request) {

	// parse request / urlencoded arguments (must be called manually)
	request.ParseForm()

	// print urlencoded form-data, etc...
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)

	// print full url request
	fmt.Println(request.Form["test_param"])

	// loop through tuple return value from request.Form
	// assign first tuple value to "key", assign the second value
	// to "val" (key, value pairs)
	for key, val := range request.Form {
		fmt.Println("key", key)
		fmt.Println("value", val)
	}

	// reply to client
	fmt.Fprintf(writer, "Hello World!")

}

func main() {

	// set router for root url requests
	http.HandleFunc("/", sayhelloWorld)
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}