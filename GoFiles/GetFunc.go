//====================================================//
//             AUTHOR: 	Brede Fritjof Klausen         //
//         UNIVERSITY: 	NTNU in Gjøvik                //
//====================================================//

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// GetValues function for getting a commit
func GetValues(URL string) Bitbucket {
	// TODO : Comment this

	payload := Bitbucket{}

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "PushInfo")

	client := http.Client{
		Timeout: time.Second * 10,
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &payload)
	if jsonErr != nil {
		if res.StatusCode == 404 {
			payload.Pagelen = -1
		} else {
			log.Fatal(jsonErr)
		}
	}

	return payload
}
