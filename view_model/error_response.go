package view_model

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"-"`
}
