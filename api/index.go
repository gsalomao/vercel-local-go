package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w,
		"<h1>Hello from Go!</h1>",
	)
}

func Main() {
	http.HandleFunc(
		"/",
		Handler,
	)

	if err := http.ListenAndServe(
		":8080",
		nil,
	); err != nil {
		panic(err)
	}
}
