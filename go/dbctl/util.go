package dbctl

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "my-secret-pw"
	PROTOCOL := "tcp(127.0.0.1:33061)"
	DBNAME := "techdojo_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

const (
	success = "success"
	failure = "failure"
	info    = "info"
)

func writeLog(level string, message ...interface{}) {
	f, err := os.OpenFile("./log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, time.Now().Format(time.RFC3339), level+"| ", message)
}
