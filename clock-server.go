package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"log"
)

type Data struct {
	Now time.Time `json:"now"`
}

func NowHandler(w http.ResponseWriter,
	r *http.Request) {
	d := Data{
		time.Now(),
	}

	bytes, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(bytes))
}

func main() {
	http.HandleFunc("/now", NowHandler)
	http.ListenAndServe(":3000", nil)
}
