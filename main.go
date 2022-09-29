package main

import (
	"fmt"
	"log"
	"net/http"
)



 // A Simple Web Server 

func formHandler(w http.ResponseWriter , r *http.Request){
	if err :=  r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succesfull\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w , "Name = %s\n",name)
	fmt.Fprintf(w, "Address = %s\n" ,address)
}

// Listener


func helloHandler(w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w , " 404 Not found" , http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w , "method is not supported" , http.StatusNotFound)
		return
	}
	fmt.Fprintf(w , "Hello!!!...")
}


func main (){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/" ,fileserver)
	http.HandleFunc("/form" , formHandler)
	http.HandleFunc("/hello" ,helloHandler)

	fmt.Printf("Server is Starting at Port 8080\n")

	if err := http.ListenAndServe(":8080" , nil); err != nil {
		log.Fatal(err)
	}

}
