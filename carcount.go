package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
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
	fmt.Println("r.Body is \n", xmlStr)


	//err := r.ParseMultipartForm(1048576)
	//if err != nil {
	//	log.Printf("Cannot ParseMultipartForm, error: %v\n", err)
	//	return
	//}
	//
	//if r.MultipartForm == nil {
	//	log.Printf("MultipartForm is null\n")
	//	return
	//}
	//
	//if r.MultipartForm.File != nil {
	//	//将图片文件名称传入解析函数，通过抓拍到的内容进行不同的处理。
	//	imageType := parseMultipartFormFile(r, r.MultipartForm.File)
	//	parseMultipartFormValue(r.MultipartForm.Value,imageType)
	//}
}

// parseMultipartFormValue解析表单数据
func parseMultipartFormValue(formValues map[string][]string,imageType string) {
	for _, values := range formValues {
		//log.Printf("Value formname: %s\n", formName)
		for _, value := range values {
			//log.Printf("      formdata[%d]: content=[%s]\n", i, value)

			m := make(map[string]string)
			_ = json.NewDecoder(strings.NewReader(value)).Decode(&m)
			//log.Printf("      Formdata[%d]: \njson=[%v]\n", i, value)

			/*var ve vehic
			j  := []byte(value)
			json.Unmarshal(j, &ve)
			fmt.Printf("测试解析%s\n 车牌：%s\n",
					ve[i].ChannelName,ve[0].CaptureResult[0].Vehicle.Property[2].Value)*/

			//这里使用的绝对路径进行数据获取，相对来说实现起来比较快。
			//目前有用数据很少，直接获取即可
			if imageType == "plateImage.jpg" {
				plateNo := gjson.Get(value, "CaptureResult.0.Vehicle.Property.2.value")
				vehicleType := gjson.Get(value, `CaptureResult.0.Vehicle.Property.#(description="vehicleType").value`)
				fmt.Printf("通道名称：%s  车牌:%v  车辆类型：%v\n",
					gjson.Get(value, "channelName"), plateNo.Value(), vehicleType.Value())
			}else if imageType == "faceImage.jpg"{
				humanAge := gjson.Get(value, `CaptureResult.0.Face.Property.#(description="age").value`)
				humanGender := gjson.Get(value, `CaptureResult.0.Face.Property.#(description="gender").value`)
				vehicleType := gjson.Get(value, `CaptureResult.0.NonMotor.Property.#(description="nonMotorType").value`)
				fmt.Printf("通道名称：%s  车辆类型：%v 性别：%s 年龄：%.0f \n",
					gjson.Get(value, "channelName"), vehicleType.Value(),
					humanGender.Value(),humanAge.Value())
			}else if imageType == "vehicleImage.jpg"{
				//
			}
		}
	}
}

// parseMultipartFormFile解析表单文件
func parseMultipartFormFile(r *http.Request , formFiles map[string][]*multipart.FileHeader) string{
	var imageType string
	for formName := range formFiles {
		// func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
		// FormFile returns the first file for the provided form key
		_, formFileHeader, _ := r.FormFile(formName)


		if formFileHeader.Filename == "plateImage.jpg" {
			imageType = formFileHeader.Filename
		}else if formFileHeader.Filename == "faceImage.jpg"{
			imageType = "faceImage.jpg"
		}else if formFileHeader.Filename == "vehicleImage.jpg" {
			if imageType != "plateImage.jpg"{
				log.Printf("File formname: %s, filename: %s, file length: %d\n",
					formName, formFileHeader.Filename, formFileHeader.Size)}
		}



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
	return imageType
}

func main() {
	http.HandleFunc("/hikcar", httpHandle)
	log.Fatal(http.ListenAndServe(":10180", nil))
}
