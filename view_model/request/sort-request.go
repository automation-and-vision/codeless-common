package request

type SortRequest struct {
	Column string `form:"sort.column"`
	Order  string `form:"sort.order"`
}
