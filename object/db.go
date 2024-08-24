package object

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var (
	Engine *xorm.Engine
)

func InitDB(dataSourceName string) {
	var err error
	Engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to create engine: %v", err)
	}

	if err = Engine.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	if exists, err := Engine.IsTableExist(new(User)); err != nil {
		log.Fatalf("Failed to check table existence: %v", err)
	} else if !exists {
		if err = Engine.Sync2(new(User)); err != nil {
			log.Fatalf("Failed to sync table: %v", err)
		}
	}
}
