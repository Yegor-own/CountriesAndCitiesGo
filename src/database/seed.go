package database

import "fmt"

func Seed() {
	countriesGen := countryFactory(10)
	citiesGen := cityFactory(50, 10)
	tx := DB.MustBegin()
	for _, v := range countriesGen {
		//fmt.Println(v)
		tx.MustExec(v)
	}
	for _, v := range citiesGen {
		//fmt.Println(v)
		tx.MustExec(v)
	}
	fmt.Printf("Seed %v countries and %v cities \n", len(countriesGen), len(citiesGen))
	err := tx.Commit()
	if err != nil {
		panic(err)
	}
	fmt.Println("Seeding complete")
}
