package main

import (
	"fmt"
	"net/http"

	//to import another folder, first letter has to be CAPS for public
	Data "./data"
	"github.com/gorilla/mux"
)

func postMethod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Data.Count++
	Data.Title[vars["title"]] = Data.Count
	fmt.Fprintf(w, "You've added the book: %s \n", vars["title"])
}
func getMethod(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "index - name\n")
	for name,index := range Data.Title{
		fmt.Fprintf(w, "%d \t- %s\n", index, name)
	}
}
func deleteMethod(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	delete(Data.Title, vars["title"])
	fmt.Fprintf(w, "You've deleted the book")
}
func main() {
	r := mux.NewRouter()

	Data.Title = make(map[string]uint32)

	r.HandleFunc("/books/{title}", postMethod).Methods("POST")
	r.HandleFunc("/books", getMethod).Methods("GET")
	r.HandleFunc("/books/{title}", deleteMethod).Methods("DELETE")

	http.ListenAndServe(":80", r)
}
