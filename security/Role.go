package security

import (
	"database/sql"
	"time"
)

//Role ...
type Role struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Data        string `json:"data"`
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

//LoadFromRow ...
func (item *Role) LoadFromRow(row *sql.Row) error {
	return row.Scan(&item.ID, &item.Description, &item.Data)
}

//LoadFromRows ...
func (item *Role) LoadFromRows(rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Description, &item.Data)
}

//GetRoles ..
func GetRoles(db *sql.DB, sql string, args ...interface{}) ([]Role, error) {
	var list []Role
	rows, err := db.Query(sql, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := Role{}
		item.LoadFromRows(rows)
		if err != nil {
			break
		}
		list = append(list, item)
	}
	return list, err
}

//GetAllRoles ..
func GetAllRoles(db *sql.DB) ([]Role, error) {
	return GetRoles(db, "select id, description, data from role")
}

//Create ...
func (item *Role) Create(db *sql.DB) error {
	_, err := db.Exec("insert into role(description, data) values (?, ?)", item.Description, item.Data)
	if err != nil {
		return err
	}
	return item.GetByDescription(db)
}

//Update ...
func (item Role) Update(db *sql.DB) error {
	_, err := db.Exec("update role set description = ?, data = ? where id = ?", item.Description, item.Data, item.ID)
	if err != nil {
		return err
	}
	userLog := UserLog{UserID: item.ID, LogDate: time.Now(), Action: "Update"}
	return userLog.Create(db)
}

//Delete ...
func (item Role) Delete(db *sql.DB) error {
	_, err := db.Exec("delete from role where id = ?", item.ID)
	return err
}

//GetRolesForUser ...
func GetRolesForUser(db *sql.DB, user User) ([]Role, error) {
	return GetRoles(db, "select r.id, r.description, r.data from userrole ur inner join role r on ur.roleid = r.id where ur.userid = ?", user.ID)
}

//GetByDescription ..
func (item *Role) GetByDescription(db *sql.DB) error {
	row := db.QueryRow("select id, description, data from role where description = ?", item.Description)
	return item.LoadFromRow(row)
}
