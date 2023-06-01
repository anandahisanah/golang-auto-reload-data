package controller

import (
	"assignment-3/database"
	"assignment-3/models"
	_ "fmt"
	_ "html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateLog(c *gin.Context) {
	db := database.GetDB()

	rand.Seed(time.Now().UnixNano())

	// generate random int
	randomIntWater := rand.Intn(100) + 1
	randomIntWind := rand.Intn(100) + 1

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

	c.HTML(http.StatusOK, "index.html", gin.H{
		"water":       Log.Water,
		"statusWater": Log.StatusWater.Name,
		"wind":        Log.Wind,
		"statusWind":  Log.StatusWind.Name,
	})
}
