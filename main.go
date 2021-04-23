package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ike-cunha/csv-parser-insert/db"
	_ "github.com/lib/pq"
)

const PORT = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/send-file", sendFile)

	server := &http.Server{
		Addr:           PORT,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Server is running on port: %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}

//Receives an file with CSV structure
func sendFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		_, writeErr := w.Write([]byte("This route only accepts POST requests"))
		if writeErr != nil {
			fmt.Println(writeErr)
			return
		}
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("data")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	//Insert data in postgres
	go db.Insert(content)

	fmt.Fprintf(w, "Uploaded file %s with success. Data is being processed.", handler.Filename)
}
