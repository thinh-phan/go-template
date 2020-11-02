package main

import (
	"net/http"

	"thinh-phan.com/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
