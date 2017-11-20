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
	http.HandleFunc("/", HandleHTML)
	http.HandleFunc("/url_is/", HandleBitbucket)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}
