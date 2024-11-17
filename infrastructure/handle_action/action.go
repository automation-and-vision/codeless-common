package handle_action

type IAction interface {
	GetType() string
	GetPayload() interface{}
}

type Action[T interface{}] struct {
	Type    string
	Payload T
}

func (t Action[T]) GetType() string {
	return t.Type
}

func (t Action[T]) GetPayload() interface{} {
	return t.Payload
}
