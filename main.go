package main

import (
	"github.com/alexandre-dos-reis/go-templ-web/components"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		components.Hello("Alexandre").Render(r.Context(), w)
	})

	http.ListenAndServe(":8080", nil)
}
