package view_model

type ErrorResponse struct {
	Err  string `json:"error"`
	Code int    `json:"-"`
}

func (r ErrorResponse) Error() string {
	return r.Err
}
