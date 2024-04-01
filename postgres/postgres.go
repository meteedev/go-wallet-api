package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DbConfiguration struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

func New() (*Postgres, error) {
	databaseSource := generateDatabaseUrl()
	//log.Println(databaseSource)
	db, err := sql.Open("postgres", databaseSource)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &Postgres{Db: db}, nil
}

func initDbConfiguration() *DbConfiguration {
	return &DbConfiguration{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB_NAME"),
		SslMode:  os.Getenv("POSTGRES_SSL_MODE"),
	}
}

func generateDatabaseUrl()string{
	dbConfig := initDbConfiguration()
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port, 
		dbConfig.User, 
		dbConfig.Password, 
		dbConfig.DbName,dbConfig.SslMode,
	)

}