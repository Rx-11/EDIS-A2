package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Rx-11/EDIS-A2/book-web-bff/config"
	"github.com/Rx-11/EDIS-A2/book-web-bff/pkg/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db            *gorm.DB
	currentDBType DBType
)

type DBType int

const (
	Postgres DBType = iota
	MySQL
	SQLite
)

type LogLevel int

const (
	LogDisable LogLevel = iota
	LogInfo
	LogWarn
	LogErr
)

func GetDB() *gorm.DB {
	return db
}

func Init(cfg config.DBConfig, dbType DBType, dbLog LogLevel) *gorm.DB {
	currentDBType = dbType
	var (
		sqlDB *sql.DB
		err   error
	)

	var dialector gorm.Dialector

	switch dbType {
	case Postgres:
		dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=require password=%s",
			cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
		dialector = postgres.Open(dsn)
	case MySQL:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
		dialector = mysql.Open(dsn)
	case SQLite:
		dsn := cfg.DBName
		if dsn == "" {
			dsn = "gorm.db"
		}
		dialector = sqlite.Open(dsn)
	default:
		log.Fatal("Unsupported DBType")
	}

	var gormLogger logger.Interface = logger.Default.LogMode(logger.Silent)
	switch dbLog {
	case LogInfo:
		gormLogger = logger.Default.LogMode(logger.Info)
	case LogWarn:
		gormLogger = logger.Default.LogMode(logger.Warn)
	case LogErr:
		gormLogger = logger.Default.LogMode(logger.Error)
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		Logger:      gormLogger,
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("GORM connection failed:", err)
	}

	sqlDB, err = db.DB()
	if err != nil {
		log.Fatal("Failed to get DB from GORM:", err)
	}

	if sqlDB != nil {
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetMaxOpenConns(35)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)
		if err := sqlDB.Ping(); err != nil {
			log.Panicln("Failed to ping database:", err)
		}
	}

	log.Println("Connected to the database!")
	return db
}

func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
	log.Println("Database connection closed.")
}

func Migrate() {
	if config.GetConfig().DbConfig.DBMigrate {
		log.Println("Migrating database...")
		if currentDBType == Postgres {
			db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
		}
		err := db.AutoMigrate(&models.Book{}, &models.User{})
		if err != nil {
			panic(err)
		}
	}
}
