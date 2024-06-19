package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("hello world\nPORT=%s\nDATABASE_URL=%s", port, databaseURL)
		io.WriteString(w, message)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
