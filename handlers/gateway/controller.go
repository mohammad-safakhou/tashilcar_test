package gateway

import (
	"context"
	"tashil/handlers"
	"tashil/models"
)

type controllerHandler struct {
	endpointHandler handlers.EndpointHandler
}

func NewControllerHandler(endpointHandler handlers.EndpointHandler) handlers.Controller {
	return &controllerHandler{endpointHandler: endpointHandler}
}

// RegisterEndpoint: registering our endpoint rules and pushing scheduling
func (e *controllerHandler) RegisterRule(ctx context.Context) error {
	_, err := e.endpointHandler.RegisterEndpoint(ctx, models.RegisterEndpoint{
		Endpoint:   models.EndpointRules{},
		Scheduling: models.Scheduling{},
	})
	if err != nil {
		return err
	}
	return nil
}
