//====================================================//
//             AUTHOR:  Brede Fritjof Klausen         //
//         UNIVERSITY:  NTNU in Gjøvik                //
//====================================================//

package main

import (
	"net/http"

	"os"
)

func main() {

	port := os.Getenv("PORT")
	http.HandleFunc("/", HandleHTML)
	http.HandleFunc("/url_is/", HandleBitbucket)
	http.ListenAndServe(":"+port, nil)
	//http.ListenAndServe("localhost:8080", nil)

}
