package tasks

import (
  "fmt"
  "github.com/chuckpreslar/gofer"
  "github.com/DavidHuie/gomigrate"
  _ "github.com/lib/pq"
  "github.com/jinzhu/gorm"
)

var Migrate = gofer.Register(gofer.Task{
  Label: "migrate",
  Description: "Migrate the database because",
  Action: func() {
    db, err := gorm.Open("postgres", "dbname=inventory sslmode=disable")

    if err != nil {
      fmt.Println(err)
    }

    //Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
    db.DB()

    // Then you could invoke `*sql.DB`'s functions with it
    db.DB().Ping()
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)

    // Disable table name's pluralization
    db.SingularTable(true)

    db.CreateTable(&Thing{})
    db.CreateTable(&Author{})
    db.CreateTable(&Publisher{})
  },
})