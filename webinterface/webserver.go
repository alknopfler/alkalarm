

package main

import (
"log"
"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Listening... :80")
	http.ListenAndServe(":80", nil)


}

