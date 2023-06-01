package main

import (
	"assignment-3/database"
	"assignment-3/models"
	"assignment-3/routes"
	_ "assignment-3/routes"
	"encoding/json"
	"fmt"
	_ "html/template"
	"io/ioutil"
	"log"
	_ "math/rand"
	_ "net/http"
	_ "time"

	_ "github.com/gin-gonic/gin"
)

func main() {
	database.StartDB()

	/*
		seed
	*/

	// truncate
	err := database.GetDB().Exec("TRUNCATE TABLE statuses CASCADE").Error
	if err != nil {
		log.Fatal("Error truncating table:", err)
	}
	// find json file
	filePath := "./database/data/statuses.json"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln("Error opening JSON file:", err)
	}

	// parse json to slice of struct
	var statuses []models.Status
	err = json.Unmarshal(file, &statuses)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	// create data
	for _, item := range statuses {
		err := database.GetDB().Create(&item).Error
		if err != nil {
			log.Fatal("Error saving to database:", err)
		}
	}

	fmt.Println("Seeding Status complete")

	// start server
	routes.StartServer().Run(":8080")
}
