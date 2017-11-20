//====================================================//
// 		   AUTHOR: 	Brede Fritjof Klausen             //
// 	   UNIVERSITY: 	NTNU in Gj�vik                    //
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

	// Declare it's a json file
	http.Header.Add(w.Header(), "Content-type", "application/json")

	// Split url for each "/"
	parts := strings.Split(r.URL.Path, "/")

	// Get length of array
	length := len(parts) - 1

	// (If the url is the right length or the length is 5 but parts[5] is blank) and parts[4] is not blank
	if (length == 4 || length == 5 && parts[5] == "") && parts[4] != "" {

		// Checks the domain name
		switch parts[2] {

		// Acceptable if the domain is bitbucket.org
		case "bitbucket.org":

			// Make api url to get json from the repository
			url := "https://api.bitbucket.org/2.0/repositories/" + parts[3] + "/" + parts[4] + "/commits"

			// Create empty struct
			info := Bitbucket{}

			// Get info from api site
			info = GetValues(url)

			// Set useful data in new struct
			thisCommit := info.Values[0]
			date := thisCommit.Date.Format("Mon. 02. January 2006 @ 15:04:05")
			showinfo := ShowInfo{thisCommit.Author.User.DisplayName, thisCommit.Author.User.UserName, thisCommit.Message, thisCommit.Links.HTML.Href, date}

			// Convert to json
			json.NewEncoder(w).Encode(showinfo)

		// Also acceptable if the domain is github.com, bur it's not implemented yet
		// TODO : Implement for github later
		case "github.com":

			// Error because it hasn't been implemented yet
			http.Error(w, "Not implemented yet", http.StatusNotImplemented)
			return

		// If the domain is neither, give error
		default:

			// Error if the domain isn't bitbucket.org
			http.Error(w, "Domain can not be '"+parts[2]+"', it has to be 'bitbucket.org' or 'github.com'", http.StatusBadRequest)
			return
		}

		// if the url is wrong
	} else {

		// Give error
		http.Error(w, "Wrong url! Format : <root>/bitbucket.org/<owner>/<repository>", http.StatusBadRequest)
	}
}

// HandleHTML handles welcome page
func HandleHTML(w http.ResponseWriter, r *http.Request) {

	// Print message to user about how to use the site
	fmt.Fprintln(w, "Hello! Welcome to my project!\n\nTo get the latest commit in json format, you have to write like this:")
	fmt.Fprintln(w, "https://www.fritjof.no/url_is/bitbucket.org/<owner>/<repository>\n\nExample:")
	fmt.Fprintln(w, "https://www.fritjof.no/url_is/bitbucket.org/Brede_F_Klausen/bitbucket-webhook\n\nStatusCode:", http.StatusOK)
}
