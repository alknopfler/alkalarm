

package main

import (
"log"
"net/http"
	"strings"
)

/*func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":80", nil)


}
*/
func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", noDirListing(http.FileServer(http.Dir("./"))))
	log.Println(http.ListenAndServe(":80", nil))
}
