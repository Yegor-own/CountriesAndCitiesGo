package database

var shema = `
DROP TABLE IF EXIST countries;
DROP TABLE IF EXIST cities;
CREATE TABLE countries (
    id int NOT NULL PRIMARY KEY,
    country varchar(255),
);
CREATE TABLE cities (
    id int NOT NULL PRIMARY KEY,
    city varchar(255),
    country_id int,
);
`

func Migrate(seed bool) {
	DB.MustExec(shema)
	if seed {
		Seed()
	}
}
