package view_model

type PageViewModel[T any] struct {
	Elements []T   `json:"elements"`
	Skip     int   `json:"skip"`
	Total    int64 `json:"total"`
}
