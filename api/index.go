package handler

import (
	"fmt"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w,
		"<h1>Hello from Go!</h1><br/>",
	)

	for _, v := range os.Environ() {

		fmt.Fprintln(w, v+"<br/>")

	}

}
