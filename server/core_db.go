package main

// import (
// 	"database/sql"
// 	"os"

// 	"github.com/sirupsen/logrus"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// // DB or Collection
// var (
// 	db_main *gorm.DB

// 	db_main_stat *sql.DB
// )

// func initDBConnection() {
// 	logrus.Printf("Starting Db Connections...")
// 	var err error
// 	db_main, err = gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
// 	if err != nil {
// 		logrus.Fatalf("Failed connect to DB main: %v", err)
// 		os.Exit(1)
// 		return
// 	}

// 	db_main_stat, err = db_main.DB()
// 	if err != nil {
// 		logrus.Fatalf("Error cannot initiate connection to DB main: %v", err)
// 		os.Exit(1)
// 		return
// 	}

// 	db_main_stat.SetMaxIdleConns(10)
// 	db_main_stat.SetMaxOpenConns(100)

// 	err = db_main_stat.Ping()
// 	if err != nil {
// 		logrus.Fatalf("Cannot ping DB main: %v", err)
// 		os.Exit(1)
// 		return
// 	}

// }

// func closeDBConnections() {
// 	logrus.Print("Closing DB Main Connection ... ")
// 	if err := db_main_stat.Close(); err != nil {
// 		logrus.Fatalf("Error on disconnection with DB Main : %v", err)
// 	}
// 	logrus.Println("Success")

// }
