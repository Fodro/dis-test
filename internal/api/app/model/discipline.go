package api

type DisciplineList struct {
	Disciplines []Discipline
}

type Discipline struct {
	ID            uint
	Title         string
	CreditUnits   float64
	AcademicHours uint32
	Prerequisites []Discipline
}
