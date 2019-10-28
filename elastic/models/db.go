package models

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"sync"
)

var dbOnce sync.Once
var DB *gorm.DB

var esOnce sync.Once
var EsClient *elasticsearch.Client

func GetDB() *gorm.DB {
	dbOnce.Do(initDB)
	return DB
}

func initDB()  {
	db, err := gorm.Open("sqlite3", "elastic.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Article{})

	DB = db
}

func GetEsClient() *elasticsearch.Client {
	esOnce.Do(initEsClient)
	return EsClient
}

func initEsClient() {
	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 1. Get cluster info
	//
	res, err := esClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	EsClient = esClient
}
