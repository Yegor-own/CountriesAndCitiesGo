package main

import (
	"countries-and-cities/database"
	"flag"
	"fmt"
)

func init() {
	migrate := flag.Bool("migrate", true, "migrate database tables")
	seed := flag.Bool("seed", true, "seed db")
	flag.Parse()
	if *migrate {
		database.Migrate(*seed)
	}
}

func main() {
	fmt.Println("hrllo")
}