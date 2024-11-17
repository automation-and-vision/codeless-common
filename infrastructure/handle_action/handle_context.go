package handle_action

type HandleContext struct {
	Request interface{}
}

type HandleResponse struct {
	Error        error
	Response     interface{}
	RowsAffected int64
}
