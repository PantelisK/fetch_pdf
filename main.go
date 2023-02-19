package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
	"errors"
	//"path"
	"log"
	"io/ioutil"
	//"net/url"
	_ "image/png"
	//"github.com/gorilla/mux"
	//"time"
	//"mimeType"
)


func getRoot(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://95.217.49.162:3000/")
	if err != nil {
		log.Fatal(err)
	}
	
	defer resp.Body.Close()
	
	// Read the response body as a byte slice
	bytes, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Fatal(err1)
	}
	
	mimeType := http.DetectContentType(bytes)
	//fmt.Println(mimeType) // image/png
if mimeType == "application/pdf" {
		fmt.Println("pdf FUNC") // image/png
		err3 := DownloadFile("dummy.pdf", "http://95.217.49.162:3000/dummy.pdf")
		if err3 != nil {
		panic(err3)

		}
		fmt.Println("Downloaded: " + "pdf")
		http.ServeFile(w, r, "./dummy.pdf")
	return
}
if mimeType == "image/png"{
		fmt.Println("png FUNC") // image/png
		err2 := DownloadFile("dummy.png", "http://95.217.49.162:3000/dummy.png")
		if err != nil {
			panic(err2)
		}
		fmt.Println("Downloaded: " + "png")
		http.ServeFile(w, r, "./dummy.png")
	return
}
fmt.Println("CORRUPT FILE ")
}




func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err

}


func main(){
	http.HandleFunc("/", getRoot)
	err1 := http.ListenAndServe(":8081", nil)
	if errors.Is(err1, http.ErrServerClosed){
		fmt.Printf("server closed")
	}else if err1 != nil{
	fmt.Printf("error starting:", err1)
	os.Exit(1)
	}



}
