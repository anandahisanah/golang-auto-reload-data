package main

import (
	"assignment-3/database"
	"assignment-3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func main() {
	database.StartDB()

	seederStatus()

	// channel
	done := make(chan bool)

	go createLog(done)

	<-done
}

func createLog(done chan<- bool) {
	for {
		db := database.GetDB()

		rand.Seed(time.Now().UnixNano())

		// generate random int
		randomIntWater := rand.Intn(20) + 1
		randomIntWind := rand.Intn(20) + 1

		// find water status
		var statusWater models.Status
		db.Order("id DESC").First(&statusWater, "code = ? AND range_start <= ?", "Water", randomIntWater)

		var statusWind models.Status
		db.Order("id DESC").First(&statusWind, "code = ? AND range_start <= ?", "Wind", randomIntWind)

		Log := models.Log{
			StatusWaterId: statusWater.Id,
			StatusWindId:  statusWind.Id,
			Water:         randomIntWater,
			Wind:          randomIntWind,
		}

		err := db.Create(&Log).Error
		if err != nil {
			log.Fatalln(err)
		}

		err = db.Preload("StatusWater").Preload("StatusWind").First(&Log).Error
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Water :", Log.Water)
		fmt.Println("Status Water :", Log.StatusWater.Name)
		fmt.Println("Wind :", Log.Wind)
		fmt.Println("Status Wind :", Log.StatusWind.Name)
		fmt.Println("\nMenunggu 15 detik...\n")

		time.Sleep(time.Second * 15)
	}
	done <- true
}

func seederStatus() {
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
}
