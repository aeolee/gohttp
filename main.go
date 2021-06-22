package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(r.Body)
	_, err := fmt.Fprintf(w, "This is HelloServer func")
	if err != nil {
		return
	}
	s, _ := ioutil.ReadAll(r.Body)
	fmt.Println("This is from HelloServer func")
	fmt.Println("r.Method is ", r.Method, " url is ", r.URL)
	fmt.Println("r.header is ", r.Header)
	fmt.Println("r.Body is \n", string(s))
	fmt.Println(time.Now())
}
func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":10080", nil))
}
