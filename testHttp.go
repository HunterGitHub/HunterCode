package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
//		http.Error(w, "sth failed", http.StatusInternalServerError)
		fmt.Println(*r)
	}

	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())
}
