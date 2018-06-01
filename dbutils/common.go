package dbutils

import "database/sql"

//CheckErr ...
func CheckErr(err error) {
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
}
