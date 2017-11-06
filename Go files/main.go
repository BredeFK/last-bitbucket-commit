package main

import (
	"net/http"
	"strings"
	"fmt"
	"time"
	"log"
	"io/ioutil"
	"encoding/json"
	"os"
)

// STRUCT FOR GETTING THE JSON FILE FROM THE API URL
type Bitbucket struct {
	Pagelen int `json:"pagelen"`
	Commit map[int]Commit
}

type Commit struct{
	Hash 	string `json:"hash"`
	Repository Repository
	Author Author
	Date 	string `json:"date"`
	Message string `json:"message"`
	Type 	string `json:"type"`
}

type Repository struct{
	Type		string `json:"type"`
	Name		string `json:"name"`
	FullName	string `json:"full_name"`
}

type Author struct{
	Raw 	string `json:"raw"`
	User User
}

type User struct{
	UserName 	string `json:"username"`
	DisplayName	string `json:"display_name"`
	Type 		string `json:"type"`
}


// FUNCTION FOR GETTING THE JSON INFORMATION FROM THE API URL
func getInfo(URL string) Bitbucket{

	client := http.Client{
		Timeout: time.Second * 2,
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

	bitbucket := Bitbucket{}
	jsonErr := json.Unmarshal(body, &bitbucket)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}



	return bitbucket
}

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
		// PRINT API URL TO USER

		// GET INFO FROM API SITE
		info := getInfo(url)
		json.NewEncoder(w).Encode(info.Pagelen)

		fmt.Fprint(w, url)
	}else{
		// ERROR IF THE DOMAIN INS'T BITBUCKET
		http.Error(w,"Domain can not be '" + parts[2] + "', it has to be 'bitbucket.org'",http.StatusMethodNotAllowed)
	}


}

func main(){
	port := os.Getenv("PORT")
	http.HandleFunc("/url_is/", HandleBitbucket)
	http.ListenAndServe(":"+port, nil)
//	http.ListenAndServe("localhost:8080", nil)
}
