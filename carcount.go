package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func CarServer(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(r.Body)
	_, err := fmt.Fprintf(w, "SUCCESSFUL!")
	if err != nil {
		checkErr(err, "远端无响应。")
	}

	fmt.Println("This is from HelloServer func")
	fmt.Println("r.Method is ", r.Method, " url is ", r.URL)
	fmt.Println("r.header is ", r.Header)

}
func main() {
	http.HandleFunc("/hikcar", HelloServer)
	log.Fatal(http.ListenAndServe(":10180", nil))
}