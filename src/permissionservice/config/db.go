package config

import (
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB Service database component
type DB struct {
	DB *gorm.DB
}

// NewDB Initial service's DB
func NewDB(c *Config) *gorm.DB {

	switch c.Database.Driver {
	case "mysql":
		return newMysqlDB(c)
	default:
		return newMysqlDB(c)
	}

}

// newMysqlDB Initial mysql db
func newMysqlDB(c *Config) *gorm.DB {

	// Database configuration
	dbConf := c.Database
	if dbConf.Database == "" {
		panic("database config is empty.")
	}

	// Database connection dsn
	dsn := dbConf.UserName + ":" +
		dbConf.Password +
		"@tcp(" + dbConf.Host + ":" + strconv.Itoa(dbConf.Port) + ")/" + dbConf.Database +
		"?charset=" + dbConf.Charset + "&parseTime=True&loc=Local"

	// Set Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // Default length of the string type field
		DisableDatetimePrecision:  true,  // Disable datetime precision, not supported on databases prior to MySQL 5.6
		DontSupportRenameIndex:    true,  // Renaming indexes is done by deleting and creating new ones.
		DontSupportRenameColumn:   true,  // Rename columns with `change`, renaming columns is not supported in databases prior to MySQL 8 and MariaDB
		SkipInitializeWithVersion: false, // Automatic configuration based on version
	}

	// New mysql with config
	newMysql := mysql.New(mysqlConfig)

	// Connect mysql
	conn, err := gorm.Open(
		newMysql,
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true, // Disable automatic foreign key creation constraints
			SkipDefaultTransaction:                   true, // Close global open transactions
		})

	if err != nil {
		panic("mysql connect failed [ERROR]=> " + err.Error())
	}

	sqlDB, _ := conn.DB()
	sqlDB.SetMaxIdleConns(dbConf.MaxIdleCons)
	sqlDB.SetMaxOpenConns(dbConf.MaxOpenCons)

	return conn

}
