package main

import (
	"database/sql"
	"fmt"
	_ "github.com/wendal/go-oci8"
	"log"
	"os"
)

func clearTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.Fatalln(err)
	}
}

func main() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
	db, err := sql.Open("oci8", "KD/123@192.168.0.9:1521/BLDB")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	defer clearTransaction(tx)
	rs, err := tx.Exec("UPDATE BLCRM.KDLYZT SET DQZT = '已签收' WHERE KDDH = 12345679")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rs.RowsAffected())
	if err := tx.Commit(); err != nil {
		// tx.Rollback() 此时处理错误，会忽略doSomthing的异常
		log.Fatalln(err)
	}
}
