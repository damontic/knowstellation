package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Serve runs the knowstellation server
func Serve(args []string) (bool, error) {
	log.Println("Starting Knowstellation Server...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return false, err
	}
	return true, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\n", r.RequestURI)
	w.Write([]byte("hello<br>"))
}
