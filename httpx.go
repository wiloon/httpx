package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("method: %s\n", req.Method)
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Printf("request body: %s\n", result)

	_, _ = fmt.Fprintf(w, "Hello there!\n")
}

func main() {
	http.HandleFunc("/metric/timers", myHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
