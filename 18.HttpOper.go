package main

import (
	"fmt"
	"net/http"
)

/*
   HTTP Oper
      Net/http is a library package used to build web applications. It has HandelFunc() function which routes the incoming request to its corresponding function.
      The ListenAndServe function is used to create a resource server which listens to the provided port.
      The function someFunc has the http.ResponceWriter and http.Request type parameter.
      It is responsible to take the incoming request and after processing the return response.
*/
func main() {
	http.HandleFunc("/", MyHandler1)
	http.HandleFunc("/Kloudone", MyHandler2)
	http.ListenAndServe(":8080", nil)
}
func MyHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
func MyHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Kloudone\n")
}
