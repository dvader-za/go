package main

import (
	"database/sql"
	"time"
)

//Contact ...
type Contact struct {
	ID         int       `json:"id"`
	Firstname  string    `json:"firstname"`
	Surname    string    `json:"surname"`
	Createdate time.Time `json:"createdate"`
	Title      string    `json:"title"`
}

func appendContact(slice []Contact, data ...Contact) []Contact {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]Contact, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func (c *Contact) load(rows *sql.Rows) {
	err := rows.Scan(&c.ID, &c.Firstname, &c.Surname, &c.Title, &c.Createdate)
	checkErr(err)
}

//LoadContacts ...
func LoadContacts(rows *sql.Rows) []Contact {
	var contacts []Contact

	for rows.Next() {
		var newContact Contact
		newContact.load(rows)
		contacts = appendContact(contacts, newContact)
	}
	return contacts
}
