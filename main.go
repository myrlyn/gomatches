// GoSluts project main.go
package main

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	//_ "github.com/glebarez/sqlite" //go native sqlite impl
)

var database *gorm.DB
var cfg *ini.File

func main() {
	dbtype := "sqlite"
	dbname := "slut.sqlite"
	fmt.Println("Hello World!")
	poolConnections := true
	cfg1, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	} else {
		cfg = cfg1
	}
	if cfg.HasSection("database") {
		if cfg.Section("database").HasKey("dbtype") {
			dbtype = cfg.Section("database").Key("dbtype").String()
		}
		if cfg.Section("database").HasKey("dbname") {
			dbname = cfg.Section("database").Key("dbname").String()
		}
		if cfg.Section("database").HasKey("poolConnections") {
			pcon, err := cfg.Section("database").Key("poolConnections").Bool()
			if err == nil {
				poolConnections = pcon
			} else {
				log.Println("CANNOT PARSE poolConnections value of " + cfg.Section("database").Key("poolConnections").String() + ", Using Default value of true")
			}
		}
	}

	configDB(dbtype, dbname)
	if poolConnections {
		poolDBConnections()
	}
	configureWebServer()
}
