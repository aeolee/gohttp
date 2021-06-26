package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// HelloServer 函数中包含有解析函数Anaxml()的调用,具体查看anaxml.go文件
func HelloServer(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(r.Body)
	/*_, err := fmt.Fprintf(w, "This is HelloServer func")
	if err != nil {
		return
	}*/
	s, _ := ioutil.ReadAll(r.Body)
	xmlStr := string(s)
	fmt.Println("This is from HelloServer func")
	fmt.Println("r.Method is ", r.Method, " url is ", r.URL)
	fmt.Println("r.header is ", r.Header)
	fmt.Println("r.Body is \n", xmlStr)

	Anaxml(xmlStr)
}
func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":10080", nil))

	Select()
}
