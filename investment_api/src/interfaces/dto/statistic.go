package dto

type Statistic struct {
	TotalAssignments           int     `json:"total_assignments"`
	TotalAssignmentSuccess     int     `json:"total_assignments_succes"`
	TotalAssignmentUnSuccess   int     `json:"total_assignments_unsuccess"`
	AverageAssignmentSuccess   float32 `json:"average_assignments_success"`
	AverageAssignmentUnSuccess float32 `json:"average_assignments_unsuccess"`
}
