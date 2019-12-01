package main

import (
	"github.com/goadesign/goa"
	"github.com/golang-test/app"
)

// MessageController implements the message resource.
type MessageController struct {
	*goa.Controller
}

// NewMessageController creates a message controller.
func NewMessageController(service *goa.Service) *MessageController {
	return &MessageController{Controller: service.NewController("MessageController")}
}

// Show runs the show action.
func (c *MessageController) Show(ctx *app.ShowMessageContext) error {
	// MessageController_Show: start_implement

	// Put your logic here

	res := &app.GoaExampleMessage{}
	return ctx.OK(res)
	// MessageController_Show: end_implement
}
