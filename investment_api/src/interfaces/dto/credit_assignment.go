package dto

type CreditAssignmentResponse struct {
	Credit300 int `json:"credit_type_300"`
	Credit500 int `json:"credit_type_500"`
	Credit700 int `json:"credit_type_700"`
}

type CreditAssignmentRequest struct {
	Investment int `json:"investment"`
}
