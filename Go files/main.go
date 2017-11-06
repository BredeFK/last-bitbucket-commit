package main

import (
	"net/http"
	"strings"
	"fmt"
	"os"
)

func getInfo(URL string){
	/*
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

	json := Json{}
	jsonErr := json.Unmarshal(body, &json)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return json
	*/


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
		// info := getInfo(url)

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
