package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

// When run, this program will make a request to a server that will sleep for  an hour.
// Consequently, the program will wait for one hour and then exit.
func main() {
	svr := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(time.Hour)
			}))
	defer svr.Close()

	fmt.Println("making a request with 2s timeout")
	var netClient = &http.Client{
		Timeout: time.Second * 2,
	}
	_, err := netClient.Get(svr.URL)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("making a request with default no timeout")
	http.Get(svr.URL)
	fmt.Println("finished request")
}
