package database

import (
	"fmt"
	"math/rand"
)

func countryFactory(n int) (res []string) {
	rand.Seed(64)
	var country = `INSERT INTO countries (country) VALUES (%v);`
	for i := 0; i < n; i++ {
		res = append(res, fmt.Sprintf(country, countries[rand.Int()]))
	}
	return
}

func cityFactory(n, maxid int) (res []string) {
	rand.Seed(64)
	var city = `INSERT INTO cities (city, country_id) VALUES (%v, %v);`
	for i := 0; i < n; i++ {
		res = append(res, fmt.Sprintf(city, cities[rand.Int()], rand.Intn(maxid)))
	}
	return
}
