package main

import "database/sql"

//Role ...
type Role struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func appendRole(slice []Role, data ...Role) []Role {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]Role, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func (item *Role) load(rows *sql.Rows) {
	err := rows.Scan(&item.ID, &item.ID, &item.Description)
	checkErr(err)
}

//LoadRoles ...
func LoadRoles(rows *sql.Rows) []Role {
	var list []Role

	for rows.Next() {
		var newObject Role
		newObject.load(rows)
		list = appendRole(list, newObject)
	}
	return list
}

//CreateRole ...
func CreateRole(db *sql.DB, role Role) Role {
	return role
}

//GetRolesForUser ...
func GetRolesForUser(db *sql.DB, user User) []Role {
	rows, err := db.Query("select r.id, r.description from userrole ur inner join role r on ur.roleid = r.id where ur.userid = ?", user.ID)
	checkErr(err)
	defer rows.Close()
	return LoadRoles(rows)
}
