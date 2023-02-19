package main

import (
    "fmt"
    //"log"
	//"io/ioutil"
	//"io"
	//"time"
    "net/http"
	//"github.com/gorilla/mux"
    "os"
)

func main(){
	
    response, err := http.Get("http://95.217.49.162:3000")
	//response, err := http.Get("https://upload.wikimedia.org/wikipedia/commons/0/0e/Tree_example_VIS.jpg")
if err != nil {
    fmt.Println(err)
    return
}
defer response.Body.Close()


	// Get the content
	contentType, err := GetFileContentType(response)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content Type: " + contentType)

}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}