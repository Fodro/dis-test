package api

// Prerequisite model info
// @Description ID of prerequisite to add
type Prerequisite struct {
	ID int `json:"id"`
}

type PrerequisiteFull struct {
	DisciplineID   int `json:"discipline_id"`
	PrerequisiteID int `json:"prerequisite_id"`
}
