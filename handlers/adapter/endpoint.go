package adapter

import (
	"context"
	"tashil/handlers"
	models2 "tashil/handlers/models"
	"tashil/models"
	"tashil/repository"
	"tashil/tasks/push"
	"time"
)

type endpointHandler struct {
	endpointRepo repository.EndpointRepository
}

func NewEndpointHandler(endpointRepo repository.EndpointRepository) handlers.EndpointHandler {
	return &endpointHandler{endpointRepo: endpointRepo}
}

// RegisterEndpoint: registering our endpoint rules and pushing scheduling
func (e *endpointHandler) RegisterEndpoint(ctx context.Context, endpoint models.RegisterEndpoint) (int, error) {
	err := e.endpointRepo.SaveEndpointRules(ctx, endpoint)
	if err != nil {
		return 0, err
	}
	// push request

	return 1, nil
}

func (e *endpointHandler) CallApi(request []byte) error {
	var endpoint models2.RegisterEndpoint
	// latest log
	err := push.HttpCall()
	if err != nil && endpoint.Endpoint.EndpointLogs[0].Status || err == nil &&
		!endpoint.Endpoint.EndpointLogs[0].Status {
		// save state
		log := false
		if err == nil {
			log = true
		}
		endpoint.Endpoint.EndpointLogs = append(endpoint.Endpoint.EndpointLogs, models2.EndpointLog{
			Status:  log,
			Created: time.Now(),
		})
		// send notification
	}
	return nil
}
