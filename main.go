package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("starting server at port 8080")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		// panic(err)
		log.Fatal(err)
	}
}

func helloHandler (w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(w, "page not found",http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(w,"method not supported !",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello from HARSH !!\n")
}

func formHandler (w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w,"POST request success !!\n")

	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(w,"Name : %s\n",name)
	fmt.Fprintf(w,"Address : %s\n",address)
}