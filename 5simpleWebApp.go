package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Go is a good langauge")
	
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Expert website was design by Me Myself and I")
}

func now(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Now day")

}


func main () {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.HandleFunc("/now/", now)
	http.ListenAndServe(":8000", nil)
	
}



