package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var portFlag = flag.String("port", "8080", "Port to listen on")

func main() {
	flag.Parse()
	addr := fmt.Sprintf("0.0.0.0:%s", *portFlag)
	fmt.Printf("Listening at %s\n", addr)
	fmt.Println(http.ListenAndServe(addr, handler{}))
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v\n", r.Method, r.URL)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Println(string(b))
	fmt.Println()
	w.WriteHeader(http.StatusOK)
}
