package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
}