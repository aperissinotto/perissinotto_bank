package dto

type ApiError struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}
