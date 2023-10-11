package api

// DisciplineList model info
// @Description List of disciplines
type DisciplineList struct {
	Disciplines []Discipline `json:"disciplines"`
}

// Discipline model info
// @Description Discipline
type Discipline struct {
	ID            uint         `json:"id"`
	Title         string       `json:"title"`
	CreditUnits   float32      `json:"credit_units"`
	AcademicHours uint32       `json:"academic_hours"`
	Prerequisites []Discipline `json:"prerequisites"`
}

// DisciplineInput model info
// @Description DisciplineInput
type DisciplineInput struct {
	Title         string  `json:"title"`
	CreditUnits   float32 `json:"credit_units"`
	AcademicHours uint32  `json:"academic_hours"`
}
