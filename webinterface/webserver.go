

package main

import (
"log"
"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":80", nil)


}

