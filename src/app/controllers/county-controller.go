package controllers

import (
	"countries-and-cities/app/models"
	"countries-and-cities/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CountryIndex(c *gin.Context) {
	rows, err := database.DB.Queryx("SELECT * FROM `countries`")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	var countries []models.Country
	for rows.Next() {
		var country models.Country
		rows.StructScan(&country)
		countries = append(countries, country)
	}

	rows, err = database.DB.Queryx("SELECT * FROM `cities`")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	for rows.Next() {
		var city models.City
		rows.StructScan(&city)
		city.Country_id++
		countries[city.Country_id-1].Cities = append(countries[city.Country_id-1].Cities, city)
	}
	//fmt.Println(countries)
	//jsonResp, err := json.Marshal(countries)

	c.JSON(http.StatusOK, countries)
}

func CountryShow(c *gin.Context) {

}

func CountryCreate(c *gin.Context) {

}

func CountryUpdate(c *gin.Context) {

}

func CountryDelete(c *gin.Context) {

}
