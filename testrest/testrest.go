package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./contacts.db")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select id, surname from contact")
	checkErr(err)
	defer rows.Close()

	loadContacts(rows)

	/*for rows.Next() {
		var surname string
		var id int
		err := rows.Scan(&id, &surname)
		checkErr(err)
		fmt.Printf("id=%d, surname=%s\n", id, surname)
	}*/
	fmt.Println("starting")
	os.Exit(0)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome!")
	})

	router.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Todo Index!")
	})
	router.HandleFunc("/todos/{todoId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		todoID := vars["todoId"]
		fmt.Fprintln(w, "Todo show:", todoID)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func checkErr(err error, args ...string) {
	if err != nil {
		fmt.Println("Error")
		fmt.Printf("%q: %s", err, args)
	}
}
