package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Demo http server")

	address := "0.0.0.0" + ":" + "5454"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is request we got")
		w.Write([]byte(`<h1 "text-align: center;">Ky Harshal patty ka he </h1>`))
	})

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("htiha")
	}

}
