package database

import "fmt"

func Migrate(seed bool) {
	schema := []string{
		"DROP TABLE IF EXISTS `countries`;",
		"DROP TABLE IF EXISTS `cities`;",
		"CREATE TABLE `countries` (id INT NOT NULL PRIMARY KEY, country CHAR(255));",
		"CREATE TABLE `cities` ( id INT NOT NULL PRIMARY KEY, country_id INT, city VARCHAR(255));",
	}

	tx := DB.MustBegin()
	for _, v := range schema {
		tx.MustExec(v)
	}
	err := tx.Commit()
	if err != nil {
		panic("Smth went wrong in transaction")
	}

	fmt.Println("Migrating complete")

	if seed {
		Seed()
	}
}
