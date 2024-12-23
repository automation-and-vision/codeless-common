package handle_action

import (
	"fmt"
	"github.com/automation-and-vision/codeless-common/view_model"
)

type IDispatcher interface {
	Dispatch(action IAction)
}

type DispatcherFuncType func(ctx *HandleContext) *HandleResponse

type StateDispatcher[A IAction] struct {
	Reducer map[string]DispatcherFuncType
	Handler *interface{}
}

func (r StateDispatcher[A]) Dispatch(a A) *HandleResponse {
	if f, exists := r.Reducer[a.GetType()]; exists {
		return f(&HandleContext{Request: a.GetPayload()})
	} else {
		return &HandleResponse{
			Error: &view_model.ErrorResponse{
				Code: 500,
				Err:  fmt.Sprintf("Action %s not found", a.GetType()),
			},
		}
	}
}
