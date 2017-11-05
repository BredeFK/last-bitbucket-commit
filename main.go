package main

import (
	"net/http"
	"strings"
	"fmt"
	"os"
)
func HandleBitbucket(w http.ResponseWriter, r *http.Request) {
	// DECLARE IT'S A JSON FILE
	http.Header.Add(w.Header(), "Content-type", "application/json")

	// SPLIT URL FOR EACH "/"
	parts := strings.Split(r.URL.Path, "/")

	// THE DOMAIN HAS TO BE BITBUCKET
	url := ""
	if parts[2] == "bitbucket.org" {

		// MAKE API URL TO GET JSON FROM THE REPO
		url = "https://api.bitbucket.org/2.0/repositories/" + parts[3] + "/" + parts[4] + "/commits"
	}else{
		http.Error(w,"Domain has to be 'bitbucket.org'",http.StatusMethodNotAllowed)
	}

	// PRINT API URL TO USER
	fmt.Fprintln(w, url)
}

func main(){
	port := os.Getenv("PORT")
	http.HandleFunc("/url_is/", HandleBitbucket)
	http.ListenAndServe(":"+port, nil)
//	http.ListenAndServe("localhost:8080", nil)
}
