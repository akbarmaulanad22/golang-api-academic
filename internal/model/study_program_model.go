package model

type StudyProgramResponse struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Level            string `json:"level"`
	Accreditation    string `json:"accreditation"`
	DurationSemester int    `json:"duration_semester"`
	FacultyID        uint   `json:"faculty_id"`
}

type CreateStudyProgramRequest struct {
	Name             string `json:"name" validate:"required"`
	Level            string `json:"level" validate:"required"`
	Accreditation    string `json:"accreditation" validate:"required"`
	DurationSemester int    `json:"duration_semester" validate:"required"`
	FacultyId        uint   `json:"faculty_id" validate:"required"`
}

type UpdateStudyProgramRequest struct {
	ID               uint   `json:"-"`
	Name             string `json:"name" validate:"required"`
	Level            string `json:"level" validate:"required"`
	Accreditation    string `json:"accreditation" validate:"required"`
	DurationSemester int    `json:"duration_semester" validate:"required"`
	FacultyId        uint   `json:"faculty_id" validate:"required"`
}

type GetStudyProgramRequest struct {
	ID uint `json:"-" validate:"required"`
}

type DeleteStudyProgramRequest struct {
	ID uint `json:"-" validate:"required"`
}
