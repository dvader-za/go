package security

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

//Create ...
func (item *User) Create(db *sql.DB, role Role) (Interaction, error) {
	var interaction Interaction
	//fmt.Printf("insert into user(username, password, name, isadmin, isactive) values (%s, %s, %s, %t, %t)\n", item.Username, item.Password, item.Name, item.IsAdmin, item.IsActive)
	_, err := db.Exec("insert into user(username, password, name, isadmin, isactive) values (?, ?, ?, ?, ?)", item.Username, item.Password, item.Name, item.IsAdmin, item.IsActive)
	if err != nil {
		return interaction, err
	}
	item.GetByUsername(db)
	//fmt.Printf("insert into user(username, password, name, isadmin, isactive) values (%d, %s, %s, %s, %t, %t)\n", item.ID, item.Username, item.Password, item.Name, item.IsAdmin, item.IsActive)
	userLog := UserLog{UserID: item.ID, LogDate: time.Now(), Action: "Create"}
	err = userLog.Create(db)
	if err != nil {
		return interaction, err
	}

	var userRole = UserRole{UserID: item.ID, RoleID: role.ID}
	err = userRole.Create(db)
	if err != nil {
		return interaction, err
	}
	interaction = Interaction{Key: RandStringBytesMaskImprSrc(10), Action: "Create", UserID: item.ID, ActionDate: time.Now(), IsActive: true, ExpireDate: time.Now().AddDate(0, 0, 1)}
	err = interaction.Create(db)
	if err != nil {
		return interaction, err
	}
	interaction.GetByKey(db)
	return interaction, err
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
	return row.Scan(&item.ID, &item.Username, &item.Password, &item.Name, &item.IsAdmin, &item.IsActive)
}

//LoadFromRows ...
func (item *User) LoadFromRows(rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Username, &item.Password, &item.Name, &item.IsAdmin, &item.IsActive)
}

//GetUsersRaw ..
func GetUsersRaw(db *sql.DB, query string, args ...interface{}) ([]User, error) {
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

	var list []User
	for rows.Next() {
		item := User{}
		err := item.LoadFromRows(rows)
		if err != nil {
			break
		}
	}

	return list, err
}

//GetAllUsers ..
func GetAllUsers(db *sql.DB) ([]User, error) {
	return GetUsersRaw(db, "select id, username, password, name, isadmin, isactive from user")
}

//GetByUsername ..
func (item *User) GetByUsername(db *sql.DB) error {
	row := db.QueryRow("select id, username, password, name, isadmin, isactive from user where username = ?", item.Username)
	return item.LoadFromRow(row)
}

//Get ..
func (item *User) Get(db *sql.DB) error {
	row := db.QueryRow("select id, username, password, name, isadmin, isactive from user where id = ?", item.ID)
	return item.LoadFromRow(row)
}

//Update ...
func (item User) Update(db *sql.DB) error {
	_, err := db.Exec("update user set password = ?, name = ?, isadmin = ?, isactive = ? where id = ?", item.Password, item.Name, item.IsAdmin, item.IsActive, item.ID)
	if err != nil {
		return err
	}
	userLog := UserLog{UserID: item.ID, LogDate: time.Now(), Action: "Update"}
	return userLog.Create(db)
}

//Delete ...
func (item User) Delete(db *sql.DB) error {
	_, err := db.Exec("delete from user where id = ?", item.ID)
	return err
}

//GetRoles ...
func (item User) GetRoles(db *sql.DB) ([]Role, error) {
	return GetRolesForUser(db, item)
}

//GetLogs ...
func (item User) GetLogs(db *sql.DB) ([]UserLog, error) {
	return GetUserLogsForUser(db, item)
}

//DeleteLogs ...
func (item *User) DeleteLogs(db *sql.DB) error {
	_, err := db.Exec("delete from userlog where userid = ?", item.ID)
	return err
}

//DeleteRoles ...
func (item *User) DeleteRoles(db *sql.DB) error {
	_, err := db.Exec("delete from userrole where userid = ?", item.ID)
	return err
}
