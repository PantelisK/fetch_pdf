package main

import (
    "fmt"
    "log"
	"time"
    "net/http"
	"github.com/gorilla/mux"
    "os"
)

func serveRandomFile(w http.ResponseWriter, r *http.Request) {

    url := "http://i.imgur.com/m1UIjW1.jpg"
    // don't worry about errors
    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }
    defer response.Body.Close()

    //open a file for writing
    file, err := os.Create("/tmp/asdf.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Use io.Copy to just dump the response body to the file. This supports huge files
    http.ServeFile(w, r, "/tmp/asdf.jpg")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Success!")
}

func main() {
	http.HandleFunc("/", serveRandomFile)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		id:= mux.Vars(r)["id"]
		fmt.Fprintf(w,id)
	})
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}