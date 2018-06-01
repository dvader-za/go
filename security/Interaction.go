package security

import (
	"database/sql"
	"dbutils"
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

//LoadFromRow ...
func (item *Interaction) LoadFromRow(row *sql.Row) error {
	err := row.Scan(&item.ID, &item.Key, &item.Action, &item.UserID, &item.ActionDate, &item.IsActive, &item.ExpireDate)
	return err
}

//LoadFromRows ...
func (item *Interaction) LoadFromRows(rows *sql.Rows) error {
	err := rows.Scan(&item.ID, &item.Key, &item.Action, &item.UserID, &item.ActionDate, &item.IsActive, &item.ExpireDate)
	return err
}

//GetInteractionsRaw ..
func GetInteractionsRaw(db *sql.DB, query string, args ...interface{}) []Interaction {
	var rows *sql.Rows
	var err error
	if len(args) > 0 {
		rows, err = db.Query(query, args)
	} else {
		rows, err = db.Query(query)
	}

	dbutils.CheckErr(err)

	var list []Interaction
	for rows.Next() {
		item := Interaction{}
		err := item.LoadFromRows(rows)
		if err != nil {
			list = append(list, item)
		}
	}

	return list
}

//Create ...
func (item Interaction) Create(db *sql.DB) {
	_, err := db.Exec("insert into interaction(key, action, userid, actiondate, isactive, expiredate) values (?, ?, ?, ?, ?, ?)", item.Key, item.Action, item.UserID, item.ActionDate, item.IsActive, item.ExpireDate)
	dbutils.CheckErr(err)
}

//GetByKey ...
func (item *Interaction) GetByKey(db *sql.DB) {
	row := db.QueryRow("select id, key, action, userid, actiondate, isactive, expiredate from interaction where key = ?", item.Key)
	item.LoadFromRow(row)
}

//Update ...
func (item Interaction) Update(db *sql.DB) {
	_, err := db.Exec("update interaction set action = ?, userid = ?, actiondate = ?, isactive = ?, expiredate = ? where id = ?", item.Key, item.Action, item.UserID, item.ActionDate, item.IsActive, item.ExpireDate, item.ID)
	dbutils.CheckErr(err)
}
