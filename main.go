package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	dtobj := time.Now()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, dtobj.Format("01-02-2006 15:04:05"))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
