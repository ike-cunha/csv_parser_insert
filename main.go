package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
		w.Write([]byte("This route only accepts POST requests"))
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("data")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	// alterar response para o usuário, pode ser usado: w.Write([]byte("essa rota somente aceita chamadas post"))
	fmt.Fprintf(w, "Uploaded File: %s", handler.Filename)
}
