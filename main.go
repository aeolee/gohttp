package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// HelloServer 函数中包含有解析函数的调用,具体查看anaxml.go文件
func HelloServer(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(r.Body)
	_, err := fmt.Fprintf(w, "SUCCESSFUL!")
	if err != nil {
		checkErr(err, "远端无响应。")
	}
	s, _ := ioutil.ReadAll(r.Body)
	xmlStr := string(s)

	//fmt.Println("This is from HelloServer func")
	//fmt.Println("r.Method is ", r.Method, " url is ", r.URL)
	//fmt.Println("r.header is ", r.Header)
	//fmt.Println("r.Body is \n", xmlStr)

	count := TimeTrigger(xmlStr)

	fmt.Printf("%s\nfrom %s to %s\n", count.channelName, count.starTime, count.endTime)
	fmt.Printf("In:%d   Leave:%d\n", count.enter, count.leave)

	Insert(*count)
	//Select()
}
func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":10080", nil))
}
