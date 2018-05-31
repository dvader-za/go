package main

import (
	"database/sql"
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

func (item *User) load(rows *sql.Rows) {
	err := rows.Scan(&item.ID, &item.Username, &item.Password, &item.Name, &item.IsAdmin, &item.IsActive)
	checkErr(err)
}

func (item *User) loadRow(row *sql.Row) {
	err := row.Scan(&item.ID, &item.Username, &item.Password, &item.Name, &item.IsAdmin, &item.IsActive)
	checkErr(err)
}

func loadUsers(rows *sql.Rows) []User {
	var list []User

	for rows.Next() {
		var newObject User
		newObject.load(rows)
		list = appendUser(list, newObject)
	}
	return list
}

//GetAllUsers ...
func GetAllUsers(db *sql.DB) []User {
	rows, err := db.Query("select id, username, password, name, isadmin, isactive from user")
	checkErr(err)
	defer rows.Close()
	return loadUsers(rows)
}

//GetUser ...
func GetUser(db *sql.DB, id int) User {
	row := db.QueryRow("select id, username, password, name, isadmin, isactive from user where id = ?", id)
	var newObject User
	newObject.loadRow(row)
	return newObject
}

//GetUserByName ...
func GetUserByName(db *sql.DB, name string) User {
	row := db.QueryRow("select id, username, password, name, isadmin, isactive from user where name = ?", name)
	var newObject User
	newObject.loadRow(row)
	return newObject
}

//CreateUser ...
func CreateUser(db *sql.DB, user User) (User, Interaction) {
	_, err := db.Exec("insert into user(username, password, name, isadmin, isactive) values (?, ?, ?, ?, ?)", user.Username, user.Password, user.Name, user.IsAdmin, user.IsActive)
	checkErr(err)
	newUser := GetUserByName(db, user.Username)
	userLog := UserLog{UserID: newUser.ID, LogDate: time.Now(), Action: "Create"}
	CreateUserLog(db, userLog)
	interaction := Interaction{Key: RandStringBytesMaskImprSrc(10), Action: "Create", UserID: newUser.ID, ActionDate: time.Now(), IsActive: true, ExpireDate: time.Now().AddDate(0, 0, 1)}
	CreateInteraction(db, interaction)
	newInteraction := GetInteractionByKey(db, interaction.Key)
	return newUser, newInteraction
}

//ForgotUserPassword ...
func ForgotUserPassword(db *sql.DB, user User) Interaction {
	userLog := UserLog{UserID: user.ID, LogDate: time.Now(), Action: "ForgotPassword"}
	CreateUserLog(db, userLog)
	interaction := Interaction{Key: RandStringBytesMaskImprSrc(10), Action: "ForgotPassword", UserID: user.ID, ActionDate: time.Now(), IsActive: true, ExpireDate: time.Now().AddDate(0, 0, 1)}
	CreateInteraction(db, interaction)
	newInteraction := GetInteractionByKey(db, interaction.Key)
	user.IsActive = false
	user.Password = ""
	ModifyUser(db, user)
	return newInteraction
}

//ModifyUser ...
func ModifyUser(db *sql.DB, user User) {
	_, err := db.Exec("update user set password = ?, name = ?, isadmin = ?, isactive = ? where id = ?", user.Password, user.Name, user.IsAdmin, user.IsActive, user.ID)
	userLog := UserLog{UserID: user.ID, LogDate: time.Now(), Action: "UpdateUser"}
	CreateUserLog(db, userLog)
	checkErr(err)
}
