package models

type Country struct {
	Id      int
	Country string
	Cities  []City
}
