package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./contacts.db")
	checkErr(err)
	defer db.Close()

	fmt.Println("starting")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><body><h1>Hello World from GO!</h1></body></html>")
	})

	router.HandleFunc("/api/contact", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("select id, firstname, surname, title, createdate from contact")
		checkErr(err)
		defer rows.Close()
		contacts := LoadContacts(rows)

		/*for _, c := range contacts {
			fmt.Printf("c=> %d %s\n", c.ID, c.Surname)
		}*/
		writeJSON(w, 200, contacts)
	})
	router.HandleFunc("/todos/{todoId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		todoID := vars["todoId"]
		fmt.Fprintln(w, "Todo show:", todoID)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func writeJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func checkErr(err error, args ...string) {
	if err != nil {
		fmt.Println("Error")
		fmt.Printf("%q: %s", err, args)
	}
}
