package main

import (
	"database/sql"
	"time"
)

//Interaction ...
type Interaction struct {
	ID         int       `json:"id"`
	Key        string    `json:"key"`
	Action     string    `json:"action"`
	UserID     int       `json:"userId"`
	ActionDate time.Time `json:"actionDate"`
	IsActive   bool      `json:"isActive"`
	ExpireDate time.Time `json:"expireDate"`
}

func appendInteraction(slice []Interaction, data ...Interaction) []Interaction {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]Interaction, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func (item *Interaction) load(rows *sql.Rows) {
	err := rows.Scan(&item.ID, &item.Key, &item.Action, &item.UserID, &item.ActionDate, &item.IsActive, &item.ExpireDate)
	checkErr(err)
}

func (item *Interaction) loadRow(row *sql.Row) {
	err := row.Scan(&item.ID, &item.Key, &item.Action, &item.UserID, &item.ActionDate, &item.IsActive, &item.ExpireDate)
	checkErr(err)
}

//LoadInteractions ...
func LoadInteractions(rows *sql.Rows) []Interaction {
	var list []Interaction

	for rows.Next() {
		var newObject Interaction
		newObject.load(rows)
		list = appendInteraction(list, newObject)
	}
	return list
}

//CreateInteraction ...
func CreateInteraction(db *sql.DB, interaction Interaction) Interaction {
	return interaction
}

//GetInteractionByKey ...
func GetInteractionByKey(db *sql.DB, key string) Interaction {
	row := db.QueryRow("select id, key, action, userid, actiondate, isactive, expiredate from interaction where key = ?", key)
	var newObject Interaction
	newObject.loadRow(row)
	return newObject
}
