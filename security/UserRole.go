package main

import "database/sql"

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

func (item *UserRole) load(rows *sql.Rows) {
	err := rows.Scan(&item.UserID, &item.RoleID)
	checkErr(err)
}

//LoadUserRoles ...
func LoadUserRoles(rows *sql.Rows) []UserRole {
	var list []UserRole

	for rows.Next() {
		var newObject UserRole
		newObject.load(rows)
		list = appendUserRole(list, newObject)
	}
	return list
}

//CreateUserRole ...
func CreateUserRole(db *sql.DB, userRole UserRole) UserRole {
	return userRole
}
