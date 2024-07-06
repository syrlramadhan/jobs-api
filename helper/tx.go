package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx){
	err := recover()
	if err != nil {
		errRollBack := tx.Rollback()
		SendPanicError(errRollBack)
		panic(err)
	}else {
		errCommit := tx.Commit()
		SendPanicError(errCommit)
	}
}