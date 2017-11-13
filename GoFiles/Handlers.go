//====================================================//
// 		   AUTHOR: 	Brede Fritjof Klausen             //
// 	   UNIVERSITY: 	NTNU in Gjøvik                    //
//====================================================//

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// HandleBitbucket handles bitbucket url and returns an api url
func HandleBitbucket(w http.ResponseWriter, r *http.Request) {
	// DECLARE IT'S A JSON FILE
	http.Header.Add(w.Header(), "Content-type", "application/json")

	// SPLIT URL FOR EACH "/"
	parts := strings.Split(r.URL.Path, "/")

	// GET LENGTH OF ARRAY
	length := len(parts) - 1

	// THE DOMAIN HAS TO BE BITBUCKET
	if length == 4 && parts[4] != "" {
		if parts[2] == "bitbucket.org" {

			// MAKE API URL TO GET JSON FROM THE REPO
			url := "https://api.bitbucket.org/2.0/repositories/" + parts[3] + "/" + parts[4] + "/commits"

			// GET INFO FROM API SITE
			info := GetValues(url)
			json.NewEncoder(w).Encode(info.Values[0]) // Print latest

		} else {
			// ERROR IF THE DOMAIN INS'T BITBUCKET
			http.Error(w, "Domain can not be '"+parts[2]+"', it has to be 'bitbucket.org'", http.StatusBadRequest)
		}
	} else {
		// ERROR IF THE USER HASN'T WRITTEN AN URL
		http.Error(w, "Wrong url! Format : <root>/bitbucket.org/<owner>/<repository>", http.StatusBadRequest)
		fmt.Fprintln(w, length)
	}
}

// HandleHTML handles welcome page
func HandleHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello! Welcome to my project!\n\nTo get the latest commit in json format, you have to write like this:")
	fmt.Fprintln(w, "https://bitbucket-commit.herokuapp.com/url_is/bitbucket.org/<owner>/<repository>\n\nStatus code:", http.StatusOK)
}
