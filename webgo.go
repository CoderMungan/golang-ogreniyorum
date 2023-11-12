package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Merhaba, Attığın get isteği: %s\n", r.URL.Path)
		fmt.Println("80 Portunda Çalışıyorum")
	})

	http.ListenAndServe(":80", nil)
}
