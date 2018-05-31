package main

import (
	"database/sql"
	"time"
)

//UserLog ...
type UserLog struct {
	ID      int       `json:"id"`
	UserID  int       `json:"userId"`
	LogDate time.Time `json:"logDate"`
	Action  string    `json:"action"`
	Data    bool      `json:"data"`
}

func appendUserLog(slice []UserLog, data ...UserLog) []UserLog {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]UserLog, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func (item *UserLog) load(rows *sql.Rows) {
	err := rows.Scan(&item.ID, &item.UserID, &item.LogDate, &item.Action, &item.Data)
	checkErr(err)
}

//LoadUserLogs ...
func LoadUserLogs(rows *sql.Rows) []UserLog {
	var list []UserLog

	for rows.Next() {
		var newObject UserLog
		newObject.load(rows)
		list = appendUserLog(list, newObject)
	}
	return list
}

//CreateUserLog ...
func CreateUserLog(db *sql.DB, userLog UserLog) UserLog {
	return userLog
}
