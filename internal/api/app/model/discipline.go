package api

type DisciplineList struct {
	Disciplines []Discipline
}

type Discipline struct {
	ID            uint
	Title         string
	CreditUnits   float32
	AcademicHours uint32
	Prerequisites []Discipline
}
