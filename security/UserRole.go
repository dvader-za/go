package security

import (
	"database/sql"
	"time"
)

//UserRole ...
type UserRole struct {
	UserID int `json:"userId"`
	RoleID int `json:"roleId"`
}

func appendUserRole(slice []UserRole, data ...UserRole) []UserRole {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]UserRole, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

//LoadFromRow ...
func (item *UserRole) LoadFromRow(row *sql.Row) error {
	return row.Scan(&item.UserID, &item.RoleID)
}

//LoadFromRows ...
func (item *UserRole) LoadFromRows(rows *sql.Rows) error {
	return rows.Scan(&item.UserID, &item.RoleID)
}

//GetUserRolesRaw ..
func GetUserRolesRaw(db *sql.DB, query string, args ...interface{}) ([]UserRole, error) {
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

	var list []UserRole
	for rows.Next() {
		item := UserRole{}
		err := item.LoadFromRows(rows)
		if err != nil {
			break
		}
		list = append(list, item)
	}

	return list, err
}

// //LoadUserRoles ...
// func LoadUserRoles(rows *sql.Rows) []UserRole {
// 	var list []UserRole

// 	for rows.Next() {
// 		var item UserRole
// 		item.LoadFromRows(rows)
// 		list = appendUserRole(list, item)
// 	}
// 	return list
// }

//Create ...
func (item *UserRole) Create(db *sql.DB) error {
	_, err := db.Exec("insert into userrole(userid, roleid) values (?, ?)", item.UserID, item.RoleID)
	if err != nil {
		return err
	}
	userLog := UserLog{UserID: item.UserID, LogDate: time.Now(), Action: "RoleLink"}
	return userLog.Create(db)
}

//Delete ...
func (item UserRole) Delete(db *sql.DB) error {
	_, err := db.Exec("delete from userrole where userid = ? and roleid = ?", item.UserID, item.RoleID)
	return err
}
