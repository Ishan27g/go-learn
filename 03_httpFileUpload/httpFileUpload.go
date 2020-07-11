package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func main() {

	port := ":8082"
	httpDir := http.Dir("./static")
	httpFileServer := http.FileServer(httpDir)

	// create callback file file upload data
	go registerCallbacks(port)

	// serve static html to display file upload form and hit above callback
	http.Handle("/", httpFileServer)
	log.Fatal(http.ListenAndServe(port, nil))

}

func registerCallbacks(port string) {

	http.HandleFunc("/upload", upload)
	http.ListenAndServe(port, nil)
}

func upload(w http.ResponseWriter, r *http.Request) {

	//parse uploaded file, max size 20 mb
	err := r.ParseMultipartForm(20 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error uploading file at " + time.Now().String())
	}
	//link to same endpoint as index.html
	F, handler, err := r.FormFile("tmpFile")
	if err != nil {
		fmt.Println("Error : ", err)
	}
	defer F.Close()

	const tmpDir string = "uploadedFiles"
	var tmpName = parseFileName(handler)

	//create local file to hold content
	localFile, err := ioutil.TempFile(tmpDir, tmpName)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	defer localFile.Close()

	//read contents of FormFile and copy to localFile
	uploadFile, err := ioutil.ReadAll(F)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	_, error := localFile.Write(uploadFile)
	if error != nil {
		fmt.Println("Error : ", err)
	}

	fmt.Fprintf(w, "Uploaded Succesfully\n")
	return
}

func parseFileName(handler *multipart.FileHeader) string {

	extension := handler.Header["Content-Type"]
	desc := handler.Header["Content-Disposition"]

	extension = strings.Split(extension[0], "/")
	names := strings.Split(desc[0], "filename=\"")

	var tmpName string = "*"
	for _, name := range names {
		if strings.Contains(name, extension[1]) {
			tmpName += name
		}
	}

	return tmpName[:len(tmpName)-1]
}
