package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var bodyText = string(body)
	fmt.Println(bodyText)
	io.WriteString(w, "Testing testing 1 2 3. "+bodyText)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("www")))
	mux.HandleFunc("/test", test)
	http.ListenAndServe(":8000", mux)
}
