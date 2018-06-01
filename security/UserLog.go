package security

import (
	"database/sql"
	"dbutils"
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

//LoadFromRow ...
func (item *UserLog) LoadFromRow(row *sql.Row) {
	err := row.Scan(&item.ID, &item.UserID, &item.LogDate, &item.Action, &item.Data)
	dbutils.CheckErr(err)
}

//LoadFromRows ...
func (item *UserLog) LoadFromRows(rows *sql.Rows) {
	err := rows.Scan(&item.ID, &item.UserID, &item.LogDate, &item.Action, &item.Data)
	dbutils.CheckErr(err)
}

//GetUserLogsRaw ..
func GetUserLogsRaw(db *sql.DB, sql string, args ...interface{}) []UserLog {
	rows, err := db.Query(sql, args)
	dbutils.CheckErr(err)

	var list []UserLog
	for rows.Next() {
		item := UserLog{}
		item.LoadFromRows(rows)
		dbutils.CheckErr(err)
		list = append(list, item)
	}
	return list
}

//Create ...
func (item UserLog) Create(db *sql.DB) {
	_, err := db.Exec("insert into userlog(userid, logdate, action, data) values (?, ?, ?, ?)", item.UserID, item.LogDate, item.Action, item.Data)
	dbutils.CheckErr(err)
}

//GetAll ...
func GetAll(db *sql.DB) []UserLog {
	return GetUserLogsRaw(db, "select id, userid, logdate, action, data from userlog")
}

//GetUserLogsForUser ...
func GetUserLogsForUser(db *sql.DB, user User) []UserLog {
	return GetUserLogsRaw(db, "select id, userid, logdate, action, data from userlog where userid = ?", user.ID)
}
