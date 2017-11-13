package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"os"
)

// Bitbucket struct
type Bitbucket struct {
	Pagelen int            `json:"pagelen"`
	Values []struct {
		Hash       string     `json:"hash"`
		Repository Repository `json:"repository"`
		Author     Author     `json:"author"`
		Date       string     `json:"date"`
		Message    string     `json:"message"`
		Type       string     `json:"type"`
	} `json:"values"`
}

// Repository struct
type Repository struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

// Author struct
type Author struct {
	Raw  string `json:"raw"`
	User User   `json:"user"`
}

// User struct
type User struct {
	UserName    string `json:"username"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
}

// FUNCTION FOR GETTING THE JSON INFORMATION FROM THE API URL

// GetCommit function for getting a commit
func GetCommit(URL string) Bitbucket {

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "PushInfo")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	payload := Bitbucket{}
	jsonErr := json.Unmarshal(body, &payload)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return payload
}

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
				info := GetCommit(url)
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

// HandleID handles bitbucket url and hash id and returns api url for the id
func HandleID(w http.ResponseWriter, r *http.Request) {
	// DECLARE IT'S A JSON FILE
	http.Header.Add(w.Header(), "Content-type", "application/json")

	// SPLIT URL FOR EACH "/"
	parts := strings.Split(r.URL.Path, "/")

	// GET LENGTH OF ARRAY
	length := len(parts) - 1

	// THE DOMAIN HAS TO BE BITBUCKET
	if parts[2] == "bitbucket.org" {

		// THE URL HAS TO HAVE AN ID
		if length == 5 {

			// THE ID HAS TO BE SOMETHING
			if parts[5] != "" {

				// MAKE API URL TO GET JSON FROM THE REPO
			//	url := "https://api.bitbucket.org/2.0/repositories/" + parts[3] + "/" + parts[4] + "/commit/" + parts[5]

				// GET INFO FROM API SITE
			//	info := GetCommit(url)

				// PRINT JSON FILE TO USER
			//	json.NewEncoder(w).Encode(info)

			} else {
				// ERROR IF THERE ISN'T AN VALID ID AT THE END OF THE URL
				http.Error(w, "You must enter an valid ID!", http.StatusBadRequest)
			}
		} else {
			// ERROR IF THERE ISN'T AN ID AT THE END OF THE URL
			http.Error(w, "You must enter an ID!", http.StatusBadRequest)
		}
	} else {
		// ERROR IF THE DOMAIN INS'T BITBUCKET
		http.Error(w, "Domain can not be '"+parts[2]+"', it has to be 'bitbucket.org'", http.StatusBadRequest)
	}
}

// HandleHTML handles welcome page
func HandleHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello! Welcome to my project!\n\nTo get the api url for your bitbucket repository, you have to write like this:")
	fmt.Fprintln(w, "https://bitbucket-commit.herokuapp.com/url_is/bitbucket.org/<owner>/<repository>")
	fmt.Fprintln(w, "\nTo get api url for a specific url, you have to write like this:")
	fmt.Fprintln(w, "https://bitbucket-commit.herokuapp.com/url_id/bitbucket.org/<owner>/<repository>/<hashID>\n\nStatus code:", http.StatusOK)
//	w.WriteHeader(http.StatusOK)
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", HandleHTML)
	http.HandleFunc("/url_is/", HandleBitbucket)
	http.HandleFunc("/url_id/", HandleID)
	http.ListenAndServe(":"+port, nil)
//	http.ListenAndServe("localhost:8080", nil)
}
