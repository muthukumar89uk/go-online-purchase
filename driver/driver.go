package driver

import (
	//user defined packages
	"online/helper"
	"online/logs"

	//Inbuild packages
	"fmt"
	"os"

	//Third-party packages
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	log := logs.Log()

	//Loading a '.env' file
	if err := helper.Config(`C:\Jackupsurya\GolangTasks\Real-Time-Tasks\RTE_Jackup\.env`); err != nil {
		log.Error.Println("Error : 'Error at loading '.env' file'")
	}
	Host := os.Getenv("HOST")
	Port := os.Getenv("PORT")
	User := os.Getenv("USER")
	Password := os.Getenv("PASSWORD")
	Dbname := os.Getenv("DBNAME")

	//create a connection to database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)
	Db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		DbConnection, err := Db.DB()
		if err != nil {
			log.Error.Println("Error : 'Invalid Database connection' ", err)
		}
		defer func() {
			fmt.Println("Database closed")
			DbConnection.Close()
		}()
		panic(err)
	}
	log.Info.Printf("Message : 'Established a successful connection to %s database!!!'\n", Dbname)
	return Db
}

func TestDbConnection() *gorm.DB {
	log := logs.Log()

	//Loading a '.env' file
	if err := helper.Config(`C:\Jackupsurya\GolangTasks\Real-Time-Tasks\RTE_Jackup\.env`); err != nil {
		log.Error.Println("Error : 'Error at loading '.env' file'")
	}

	Host := os.Getenv("HOST")
	Port := os.Getenv("PORT")
	User := os.Getenv("USER")
	Password := os.Getenv("PASSWORD")
	Dbname := os.Getenv("TEST_DBNAME")

	// create a connection to database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)
	Db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		DbConnection, err := Db.DB()
		if err != nil {
			log.Error.Println("Error : 'Invalid Database connection' ", err)
		}
		defer func() {
			fmt.Println("Database closed")
			DbConnection.Close()
		}()
		panic(err)
	}
	log.Info.Printf("Message : 'Established a successful connection to %s database!!!'\n", Dbname)
	return Db
}
