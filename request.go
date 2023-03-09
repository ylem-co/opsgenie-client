package opsgenieclient

type CreateAlertRequest struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}
