package database

import "fmt"

func Seed() {
	countriesGen := countryFactory(10)
	citiesGen := cityFactory(50, 9)
	tx := DB.MustBegin()
	for _, v := range countriesGen {
		//fmt.Println(v)
		tx.MustExec(v)
	}
	for _, v := range citiesGen {
		//fmt.Println(v)
		tx.MustExec(v)
	}
	err := tx.Commit()
	if err != nil {
		panic(err)
	}
	fmt.Println("Seeding complete")
}
