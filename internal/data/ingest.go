package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/illusioniststg/hopeline/internal/domain"
)

var Survivors map[string]domain.Survivor

func ReadSurvivorsFromCSV(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Print("fix ME!", err.Error())
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1 // 🔥 IMPORTANT FIX

	line := 0
	Survivors = make(map[string]domain.Survivor)

	for {
		line++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Skipping line %d due to error: %v\n", line, err)
			continue
		}
		var survivor domain.Survivor
		survivor.FirstName = record[0]
		survivor.LastName = record[1]
		survivor.DOB = record[2]
		survivor.Gender = record[3]
		survivor.Phone = record[4]
		survivor.Village = record[5]
		survivor.District = record[6]
		survivor.State = record[7]
		survivor.IsInjured = false
		survivor.IsMissing = false
		survivor.ShelterID = -1
		survivor.BoatID = -1

		locs := record[8]
		pairs := strings.Split(locs, "|")
		var locations []domain.Location
		for _, p := range pairs {
			coords := strings.Split(p, ", ")
			lat, err1 := strconv.ParseFloat(coords[0], 64)
			lon, err2 := strconv.ParseFloat(coords[1], 64)
			if err1 != nil || err2 != nil {
				continue
			}

			locations = append(locations, domain.Location{
				Latitude:  lat,
				Longitude: lon,
			})
		}
		survivor.Locs = locations
		Survivors[survivor.Phone] = survivor
	}

	return nil
}

func GetSurvivors(c *gin.Context) {
	filePath := "internal/data/basedata.csv"
	err := ReadSurvivorsFromCSV(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Survivors)
}
