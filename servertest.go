package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)
var success = "success"
func main() {

	//log.Println("server started")
	//	//http.HandleFunc("/", handler)
	//	//log.Println(http.ListenAndServe("localhost:6060", nil))

	test := false
	if !test {
		fmt.Println("reversed.")
	}
	fmt.Println(!test)
	fmt.Println(test)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"success"}`))
}