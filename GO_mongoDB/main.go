package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"main.go/includes"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
	includes.Connect()
	CreateReadisConnection()
	// handel.HandleRequest()
	readExcelFile("Data-for-Practice.xlsx")
}
func CreateReadisConnection() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})
}
func readExcelFile(filePath string) {
	// Open the Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
	}
	defer f.Close()

	// Specify the sheet and iterate over the rows
	sheetName := "Problem" // Replace with your sheet name
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Failed to get rows: %v", err)
	}
	var newDocuments []interface{}
	var backupDocuments []interface{}
	headers := rows[0]
	batchSize := 10
	backupDocumentsBatchSize := 1000
	for _, row := range rows[1:] {
		rowString := strings.Join(row, "|")
		hash := createSHA256Hash(rowString)

		val, err := rdb.Get(ctx, hash).Result()
		if err == redis.Nil {
			fmt.Println("Key does not exist, inserting new value")
			// Key does not exist, so insert it
			err := rdb.Set(ctx, hash, rowString, 0).Err()
			if err != nil {
				log.Fatalf("Failed to set in Redis: %v", err)
			}

			document := bson.D{}
			for i, cell := range row {
				if i < len(headers) {
					document = append(document, bson.E{Key: headers[i], Value: cell})
				}
			}
			newDocuments = append(newDocuments, document)

			if len(newDocuments) >= batchSize {
				_, err = includes.DB.Database("test").Collection("new").InsertMany(context.Background(), newDocuments)
				if err != nil {
					log.Println(err)
				}
				newDocuments = nil
			}
			// -----------------backup----------------
			backpackDocument := bson.D{{"redisKey", hash}, {"rowString", rowString}}
			for i, cell := range row {
				if i < len(headers) {
					backpackDocument = append(backpackDocument, bson.E{Key: headers[i], Value: cell})
				}
			}
			backupDocuments = append(backupDocuments, backpackDocument)

			if len(backupDocuments) >= backupDocumentsBatchSize {
				_, err = includes.DB.Database("test").Collection("new2").InsertMany(context.Background(), backupDocuments)
				if err != nil {
					log.Println(err)
				}
				backupDocuments = nil
			}
		} else if err != nil {
			log.Fatalf("Error retrieving key from Redis: %v", err)
		} else {
			fmt.Println("Key exists, value: ", val)
		}

	}
	// Insert any remaining documents
	if len(newDocuments) > 0 {
		_, err = includes.DB.Database("test").Collection("new").InsertMany(context.Background(), newDocuments)
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(backupDocuments) > 0 {
		_, err = includes.DB.Database("test").Collection("new2").InsertMany(context.Background(), backupDocuments)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createSHA256Hash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
