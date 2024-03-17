package main

import "net/http"
import "github.com/antelman107/vercel_local_go/api"

func main() {
	http.HandleFunc(
		"/",
		handler.Handler,
	)

	if err := http.ListenAndServe(
		":8080",
		nil,
	); err != nil {
		panic(err)
	}
}
