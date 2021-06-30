package main

import (
	"encoding/json"
	"github.com/aeolee/cron"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

func httpHandle(w http.ResponseWriter, r *http.Request) {
	// http.Request has a member MultipartForm, it's defined as:
	// MultipartForm *multipart.Form
	// type Form struct {
	//    Value map[string][]string
	//    File  map[string][]*FileHeader
	// }

	err := r.ParseMultipartForm(1048576)
	if err != nil {
		log.Printf("Cannot ParseMultipartForm, error: %v\n", err)
		return
	}

	if r.MultipartForm == nil {
		log.Printf("MultipartForm is null\n")
		return
	}

	/*data,err := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	if r.MultipartForm.Value != nil{
		json.Unmarshal([]byte(data),&m)
		fmt.Println(m["result"].(map[string]interface{})["A"].
			([]interface{})[3].(map[string]interface{})["name"])
	}*/

	if r.MultipartForm.Value != nil {
		parseMultipartFormValue(r.MultipartForm.Value)
	}

	/*if r.MultipartForm.File != nil {
		parseMultipartFormFile(r, r.MultipartForm.File)
	}*/
}

// 解析表单数据
func parseMultipartFormValue(formValues map[string][]string) {
	for formName, values := range formValues {
		log.Printf("Value formname: %s\n", formName)
		for i, value := range values {
			//log.Printf("      formdata[%d]: content=[%s]\n", i, value)

			m := make(map[string]string)
			_ = json.NewDecoder(strings.NewReader(value)).Decode(&m)
			log.Printf("      Formdata[%d]: json=[%v]\n", i, value)
		}
	}
	return
}

// 解析表单文件
func parseMultipartFormFile(r *http.Request, formFiles map[string][]*multipart.FileHeader) {
	for formName := range formFiles {
		// func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
		// FormFile returns the first file for the provided form key
		_, formFileHeader, _ := r.FormFile(formName)

		log.Printf("File formname: %s, filename: %s, file length: %d\n",
			formName, formFileHeader.Filename, formFileHeader.Size)
		/*
			if strings.HasSuffix(formFileHeader.Filename, ".zip") {
				zipReader, _ := zip.NewReader(formFile, formFileHeader.Size)
				for i, zipMember := range zipReader.File {
					f, _ := zipMember.Open()
					defer f.Close()

					if zipMember.FileInfo().IsDir() {
						log.Printf("     formfile[%d]: filename=[%s], ISDIR\n", i, zipMember.Name)
					} else {
						buf, _ := ioutil.ReadAll(f)
						log.Printf("     formfile[%d]: filename=[%s], size=%d, content=[%s]\n", i, zipMember.Name, len(buf), strings.TrimSuffix(string(buf), "\n"))
					}
				}
			} else {
				var b bytes.Buffer
				_, _ = io.Copy(&b, formFile)
				log.Printf("     formfile: content=[%s]\n", strings.TrimSuffix(b.String(), "\n"))
			}*/
	}
}

func main() {
	http.HandleFunc("/hikcar", httpHandle)
	log.Fatal(http.ListenAndServe(":10180", nil))
	c := cron.New()
	c.Start()
}
