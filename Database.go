// Database
package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func configDB(dbtype string, dbname string) {
	log.Println(dbtype)
	log.Println(dbname)

}

func configDBsqlite(dbname string) {
	database, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = database
}

func configDBmysql(dbname string) {
	username := "slut"
	pw := "SacredSluts"
	protocol := "tcp"
	host := "localhost"
	port := "3306"
	uriparams := "charset=utf8"
	defaultStringSize := uint(256)
	disableDateTimePrecision := false
	dontSupportRenameIndex := true
	dontSupportRenameColumn := true
	skipInitializeWithVersion := true
	if cfg.Section("database").HasKey("username") {
		username = cfg.Section("database").Key("username").String()
	}
	if cfg.Section("database").HasKey("password") {
		pw = cfg.Section("database").Key("password").String()
	}
	if cfg.Section("database").HasKey("protocol") {
		protocol = cfg.Section("database").Key("protocol").String()
	}
	if cfg.Section("database").HasKey("host") {
		host = cfg.Section("database").Key("host").String()
	}
	if cfg.Section("database").HasKey("port") {
		port = cfg.Section("database").Key("port").String()
	}
	if cfg.Section("database").HasKey("uriparams") {
		uriparams = cfg.Section("database").Key("uriparams").String()
	}
	if cfg.Section("database").HasKey("defaultStringSize") {
		dfs, err := cfg.Section("database").Key("defaultStringSize").Uint()
		if err == nil {
			defaultStringSize = dfs
		} else {
			log.Println("CANNOT PARSE defaultStringSize value of " + cfg.Section("database").Key("defaultStringSize").String() + ", Using Default value of 256")
		}
	}
	if cfg.Section("database").HasKey("disableDateTimePrecision") {
		dtp, err := cfg.Section("database").Key("disableDateTimePrecision").Bool()
		if err == nil {
			disableDateTimePrecision = dtp
		} else {
			log.Println("CANNOT PARSE disableDateTimePrecision value of " + cfg.Section("database").Key("disableDateTimePrecision").String() + ", Using Default value of false")
		}
	}
	if cfg.Section("database").HasKey("dontSupportRenameIndex") {
		drni, err := cfg.Section("database").Key("dontSupportRenameIndex").Bool()
		if err == nil {
			dontSupportRenameIndex = drni
		} else {
			log.Println("CANNOT PARSE disableDateTimePrecision value of " + cfg.Section("database").Key("dontSupportRenameIndex").String() + ", Using Default value of true")
		}
	}
	if cfg.Section("database").HasKey("dontSupportRenameColumn") {
		drnc, err := cfg.Section("database").Key("dontSupportRenameColumn").Bool()
		if err == nil {
			dontSupportRenameColumn = drnc
		} else {
			log.Println("CANNOT PARSE dontSupportRenameColumn value of " + cfg.Section("database").Key("dontSupportRenameColumn").String() + ", Using Default value of true")
		}
	}
	if cfg.Section("database").HasKey("skipInitializeWithVersion") {
		skipiwv, err := cfg.Section("database").Key("skipInitializeWithVersion").Bool()
		if err == nil {
			skipInitializeWithVersion = skipiwv
		} else {
			log.Println("CANNOT PARSE skipInitializeWithVersion value of " + cfg.Section("database").Key("skipInitializeWithVersion").String() + ", Using Default value of true")
		}
	}
	dsn := username + ":" + pw + "@" + protocol + "(" + host + ":" + port + ")/" + dbname + "?" + uriparams
	db1, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         defaultStringSize,
			DisableDatetimePrecision:  disableDateTimePrecision,
			DontSupportRenameIndex:    dontSupportRenameIndex,
			DontSupportRenameColumn:   dontSupportRenameColumn,
			SkipInitializeWithVersion: skipInitializeWithVersion}),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = db1

}
func poolDBConnections() {
	sqlDB, err := database.DB()
	if err != nil {
		panic(err)
	}
	maxIdleConns := 10
	maxOpenConns := 100
	maxConnOpenMinutes := uint(60)
	if cfg.Section("database").HasKey("maxIdleConnections") {
		mic, err := cfg.Section("database").Key("maxIdleConnections").Int()
		if err == nil {
			maxIdleConns = mic
		} else {
			log.Println("CANNOT PARSE maxIdleConns value of " + cfg.Section("database").Key("maxIdleConns").String() + ", Using Default value of 10")
		}
	}
	if cfg.Section("database").HasKey("maxOpenConnections") {
		moc, err := cfg.Section("database").Key("maxOpenConnections").Int()
		if err == nil {
			maxOpenConns = moc
		} else {
			log.Println("CANNOT PARSE maxOpenConns value of " + cfg.Section("database").Key("maxOpenConns").String() + ", Using Default value of 100")
		}
	}

	if cfg.Section("database").HasKey("maxConnOpenMinutes") {
		mocmin, err := cfg.Section("database").Key("maxConnOpenMinutes").Uint()
		if err == nil {
			maxConnOpenMinutes = mocmin
		} else {
			log.Println("CANNOT PARSE maxConnOpenMinutes value of " + cfg.Section("database").Key("maxConnOpenMinutes").String() + ", Using Default value of 60")
		}
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(maxConnOpenMinutes))

}
