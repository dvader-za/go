package security

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

//LoadFromRow ...
func (item *Interaction) LoadFromRow(row *sql.Row) error {
	return row.Scan(&item.ID, &item.Key, &item.Action, &item.UserID, &item.ActionDate, &item.IsActive, &item.ExpireDate)
}

//LoadFromRows ...
func (item *Interaction) LoadFromRows(rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Key, &item.Action, &item.UserID, &item.ActionDate, &item.IsActive, &item.ExpireDate)
}

//GetInteractionsRaw ..
func GetInteractionsRaw(db *sql.DB, query string, args ...interface{}) ([]Interaction, error) {
	var rows *sql.Rows
	var err error
	if len(args) > 0 {
		rows, err = db.Query(query, args)
	} else {
		rows, err = db.Query(query)
	}

	if err != nil {
		return nil, err
	}

	var list []Interaction
	for rows.Next() {
		item := Interaction{}
		err := item.LoadFromRows(rows)
		if err != nil {
			break
		}
		list = append(list, item)
	}

	return list, err
}

//Create ...
func (item *Interaction) Create(db *sql.DB) error {
	_, err := db.Exec("insert into interaction(key, action, userid, actiondate, isactive, expiredate) values (?, ?, ?, ?, ?, ?)", item.Key, item.Action, item.UserID, item.ActionDate, item.IsActive, item.ExpireDate)
	if err != nil {
		return err
	}
	return item.GetByKey(db)
}

//GetByKey ...
func (item *Interaction) GetByKey(db *sql.DB) error {
	row := db.QueryRow("select id, key, action, userid, actiondate, isactive, expiredate from interaction where key = ?", item.Key)
	return item.LoadFromRow(row)
}

//Update ...
func (item Interaction) Update(db *sql.DB) error {
	_, err := db.Exec("update interaction set action = ?, userid = ?, actiondate = ?, isactive = ?, expiredate = ? where id = ?", item.Key, item.Action, item.UserID, item.ActionDate, item.IsActive, item.ExpireDate, item.ID)
	return err
}

//Delete ...
func (item Interaction) Delete(db *sql.DB) error {
	_, err := db.Exec("delete from interaction where id = ?", item.ID)
	return err
}
