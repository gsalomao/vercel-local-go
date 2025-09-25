package handler

import (
	"fmt"
	"net/http"
	"os"
	"slices"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w,
		"<h1>Hello from Go!</h1><br/>",
	)

	fmt.Printf("Hello from Go!\n")

	envs := os.Environ()
	slices.Sort(envs)
	for _, v := range envs {
		fmt.Printf("%v\n", v)
	}
}
