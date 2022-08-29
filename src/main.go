package main

import (
	"countries-and-cities/database"
	"flag"
	"fmt"
)

var db_interaction bool

func init() {
	db_interaction = false
	migrate := flag.Bool("migrate", false, "migrate database tables")
	seed := flag.Bool("seed", false, "seed db")
	flag.Parse()
	if *migrate {
		fmt.Println("Migrating in progress")
		database.Migrate(*seed)
		db_interaction = true
	}
	if *seed && !*migrate {
		fmt.Println("Seeding without migrating")
		database.Seed()
		db_interaction = true
	}
}

func main() {
	if !db_interaction {
		api()
	}
}
