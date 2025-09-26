package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func hello(rw http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(rw, "Hello")
}

func attack(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "U r under attack :)")

}

func launchAttack(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "U will be hacked :)")

}

func LaunchServer(port int) {

	fmt.Printf("Server launch on port %d\n", port)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/launchAttack", launchAttack)
	http.HandleFunc("/attack", attack)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
