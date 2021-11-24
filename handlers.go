package main
//	testing package
//	mock server
// sql injection

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	//"io/ioutil"
	//"encoding/json"
	"log"
)



func handler(w http.ResponseWriter, r *http.Request)	{
	//	TODO: Do smth, maybe description of valid requests
	log.Println("signal recieved")
	fmt.Fprintf(w, "Hello there!\n") // <-- works just fine
	w.WriteHeader(http.StatusAccepted)
}

// methods
func	getHandler(w http.ResponseWriter, r *http.Request)	{

	// switch r.Method
	// case GET --> get func
	// case POST --> post func
	log.Println("signal recieved")
	//	TODO: Combine all rows in database
	fmt.Fprintf(w, "Get Handler\n") // <-- works just fine
	w.WriteHeader(http.StatusAccepted)
}

func	getUniqueHandler(w http.ResponseWriter, r *http.Request)	{
	log.Println("signal recieved")
	//fmt.Fprintf(w, "Unique Handler\n") // <-- works just fine
	//fmt.Fprintf(w, "%s\n", r.Method) // <-- works just fine
	//fmt.Fprintf(w, "Request: Host: |%s| Path: |%s|\n", r.Host, r.URL.Path)
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/fact/"))
	if err != nil	{
		fmt.Fprintf(w, "Invalid Index\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Request for fact No %d\n", id) // <-- works just fine
	//	TODO: go for Query and, if found index in database, return listing
	//err = conn.QueryRow("select id, name from users where id = ?", 1).Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusFound)
}

func	generalHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404: not found.", http.StatusNotFound)
		return
	}
	log.Println("signal recieved")
	fmt.Fprintf(w, "General Handler\n") // <-- works just fine
	w.WriteHeader(http.StatusAccepted)
}

func testHandlers()	{


	sm := http.NewServeMux()
	// Let gv{1,2,3} handle routes{1,2,3} respectively
	sm.HandleFunc("/", generalHandler)
	sm.HandleFunc("/fact", getHandler)
	sm.HandleFunc("/fact/", getUniqueHandler)		// <-- regex for id (num)
	// TODO: default handler for 404
	//sm.HandleFunc("", getUniqueHandler)
	if err := http.ListenAndServe(":8080", sm); err != nil {
		log.Fatal(err)
	}

}
func initRequestHandling()	{

	testHandlers()
}
