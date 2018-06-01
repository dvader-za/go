package security

import (
	"database/sql"
	"dbutils"
	"fmt"
	"time"
)

//User ...
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	IsAdmin  bool   `json:"isAdmin"`
	IsActive bool   `json:"isActive"`
}

func appendUser(slice []User, data ...User) []User {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]User, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

//Create ...
func (item *User) Create(db *sql.DB) Interaction {
	fmt.Printf("insert into user(username, password, name, isadmin, isactive) values (?, ?, ?, ?, ?)\n", item.Username, item.Password, item.Name, item.IsAdmin, item.IsActive)
	_, err := db.Exec("insert into user(username, password, name, isadmin, isactive) values (?, ?, ?, ?, ?)", item.Username, item.Password, item.Name, item.IsAdmin, item.IsActive)
	dbutils.CheckErr(err)
	item.GetByUsername(db)
	fmt.Printf("insert into user(username, password, name, isadmin, isactive) values (?, ?, ?, ?, ?, ?)\n", item.ID, item.Username, item.Password, item.Name, item.IsAdmin, item.IsActive)
	userLog := UserLog{UserID: item.ID, LogDate: time.Now(), Action: "Create"}
	userLog.Create(db)
	interaction := Interaction{Key: RandStringBytesMaskImprSrc(10), Action: "Create", UserID: item.ID, ActionDate: time.Now(), IsActive: true, ExpireDate: time.Now().AddDate(0, 0, 1)}
	interaction.Create(db)
	interaction.GetByKey(db)
	return interaction
}

//ForgotPassword ...
func (item *User) ForgotPassword(db *sql.DB) Interaction {
	userLog := UserLog{UserID: item.ID, LogDate: time.Now(), Action: "ForgotPassword"}
	userLog.Create(db)
	interaction := Interaction{Key: RandStringBytesMaskImprSrc(10), Action: "ForgotPassword", UserID: item.ID, ActionDate: time.Now(), IsActive: true, ExpireDate: time.Now().AddDate(0, 0, 1)}
	interaction.Create(db)
	interaction.GetByKey(db)
	item.IsActive = false
	item.Password = ""
	item.Update(db)
	return interaction
}

//LoadFromRow ...
func (item *User) LoadFromRow(row *sql.Row) error {
	err := row.Scan(&item.ID, &item.Username, &item.Password, &item.Name, &item.IsAdmin, &item.IsActive)
	return err
}

//LoadFromRows ...
func (item *User) LoadFromRows(rows *sql.Rows) error {
	err := rows.Scan(&item.ID, &item.Username, &item.Password, &item.Name, &item.IsAdmin, &item.IsActive)
	return err
}

//GetUsersRaw ..
func GetUsersRaw(db *sql.DB, query string, args ...interface{}) []User {
	var rows *sql.Rows
	var err error
	if len(args) > 0 {
		rows, err = db.Query(query, args)
	} else {
		rows, err = db.Query(query)
	}

	dbutils.CheckErr(err)

	var list []User
	for rows.Next() {
		item := User{}
		err := item.LoadFromRows(rows)
		if err != nil {
			list = append(list, item)
		}
	}

	return list
}

//GetAllUsers ..
func GetAllUsers(db *sql.DB) []User {
	return GetUsersRaw(db, "select id, username, password, name, isadmin, isactive from user")
}

//GetByUsername ..
func (item *User) GetByUsername(db *sql.DB) {
	row := db.QueryRow("select id, username, password, name, isadmin, isactive from user where username = ?", item.Username)
	item.LoadFromRow(row)
	fmt.Printf("\nid %d username %s\n", item.ID, item.Username)
}

//Get ..
func (item *User) Get(db *sql.DB) {
	row := db.QueryRow("select id, username, password, name, isadmin, isactive from user where id = ?", item.ID)
	item.LoadFromRow(row)
}

//Update ...
func (item User) Update(db *sql.DB) {
	_, err := db.Exec("update user set password = ?, name = ?, isadmin = ?, isactive = ? where id = ?", item.Password, item.Name, item.IsAdmin, item.IsActive, item.ID)
	dbutils.CheckErr(err)
	userLog := UserLog{UserID: item.ID, LogDate: time.Now(), Action: "Update"}
	userLog.Create(db)
}

//Delete ...
func (item User) Delete(db *sql.DB) {
	_, err := db.Exec("delete from user where id = ?", item.ID)
	dbutils.CheckErr(err)
}

//GetRoles ...
func (item User) GetRoles(db *sql.DB) []Role {
	return GetRolesForUser(db, item)
}

//GetLogs ...
func (item User) GetLogs(db *sql.DB) []UserLog {
	return GetUserLogsForUser(db, item)
}
