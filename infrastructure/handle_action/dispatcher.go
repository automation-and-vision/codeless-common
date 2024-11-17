package handle_action

import (
	"fmt"
	"github.com/automation-and-vision/codeless-common/view_model"
)

type IDispatcher interface {
	Dispatch(action IAction)
}

type dispatcherFuncType func(ctx *HandleContext) *HandleResponse

type WorkspaceStateDispatcher[A IAction] struct {
	Reducer map[string]dispatcherFuncType
	Handler *interface{}
}

func (r WorkspaceStateDispatcher[A]) Dispatch(a A) *HandleResponse {
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
