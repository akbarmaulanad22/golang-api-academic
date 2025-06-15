package model

type ListScheduleRequest struct {
	UserID uint `json:"user_id"`
}

type ScheduleResponse struct {
	Course    string `json:"course"`
	Lecturer  string `json:"lecturer"`
	Classroom string `json:"classroom"`
	StartAt   string `json:"start_at"`
	EndAt     string `json:"end_at"`
}
