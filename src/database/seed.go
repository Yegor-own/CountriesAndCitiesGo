package database

var (
	countries = []string{"Russia", "USA", "Indonesia", "Poland", "Germany", "France", "UK"}
	cities    = []string{"Moscow", "Brno", "Manchester", "Pevek", "Leningrad"}
)

func Seed() {
	countryFactory(10)
	cityFactory(50, 9)
}
