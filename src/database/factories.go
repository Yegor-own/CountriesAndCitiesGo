package database

import (
	"fmt"
	"math/rand"
)

var (
	countries = []string{"Russia", "USA", "Indonesia", "Poland", "Germany", "France", "UK"}
	cities    = []string{"Moscow", "Brno", "Manchester", "Pevek", "Leningrad"}
)

func countryFactory(n int) (res []string) {
	rand.Seed(64)
	var country = "INSERT INTO `countries` (country) VALUES ('%v');"
	for i := 0; i < n; i++ {
		res = append(res, fmt.Sprintf(country, countries[rand.Intn(7)]))
	}
	return
}

func cityFactory(n, maxid int) (res []string) {
	rand.Seed(64)
	var city = "INSERT INTO `cities` (city, country_id) VALUES ('%v', %v);"
	for i := 0; i < n; i++ {
		res = append(res, fmt.Sprintf(city, cities[rand.Intn(5)], rand.Intn(maxid)))
	}
	return
}
