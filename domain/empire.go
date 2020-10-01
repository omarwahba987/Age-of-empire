package domain

type empireResponseStruct struct {
	Id                 int
	Name               string
	Expansion          string
	Army_type          string
	Unique_unit        []string
	Unique_tech        []string
	Team_bonus         string
	Civilization_bonus []string
}
