package auths

type ReportRequest struct {
	Month string `json:"month"`
	Limit string `json:"limit"`
	Page  string `json:"page"`
}

type MonthlyReport struct {
	ID           string `json:"guid"`
	MerchantName string `json:"merchant+name"`
}
