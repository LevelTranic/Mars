package database

import (
	"log/slog"
	"os"
	"time"

	"github.com/3JoB/ulib/litefmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"Mars/database/shared"
	"Mars/database/v2/schemas"
	"Mars/shared/configure"
)

func Database() {
	var driver gorm.Dialector
	switch configure.Get().Database.Driver {
	case "sqlite":
		driver = sqlite.Open(configure.Get().Database.Sqlite.Name)
	case "mysql":
		// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
		dsn := litefmt.Sprint(
			configure.Get().Database.Mysql.User,
			":",
			configure.Get().Database.Mysql.Pass,
			"@tcp(",
			configure.Get().Database.Mysql.Addr,
			")/",
			configure.Get().Database.Mysql.DBName,
			"?charset=",
			configure.Get().Database.Mysql.Charset,
			"&parseTime=True&loc=Local")
		mysql.New(mysql.Config{
			DSN:                       dsn,
			DisableDatetimePrecision:  configure.Get().Database.Mysql.DisableDatetimePrecision,
			DontSupportRenameIndex:    configure.Get().Database.Mysql.DontSupportRenameIndex,
			DontSupportRenameColumn:   configure.Get().Database.Mysql.DontSupportRenameColumn,
			SkipInitializeWithVersion: configure.Get().Database.Mysql.SkipInitializeWithVersion,
		})
		driver = mysql.Open(dsn)
	default:
		slog.Error("database driver not support")
		os.Exit(1)
	}
	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&schemas.Build{}, &schemas.Project{}, &schemas.Version{}, &schemas.VersionFamily{}); err != nil {
		panic(err)
	}
	shared.DB = db
	{
		d, err := shared.DB.DB()
		if err != nil {
			panic(err)
		}
		d.SetMaxIdleConns(configure.Get().Database.Global.MaxIdleConns)
		d.SetConnMaxLifetime(time.Duration(configure.Get().Database.Global.ConnMaxLifetime) * time.Second)
		d.SetMaxOpenConns(configure.Get().Database.Global.MaxOpenConns)
		d.SetConnMaxIdleTime(time.Duration(configure.Get().Database.Global.ConnMaxIdleTime) * time.Second)
	}
}
