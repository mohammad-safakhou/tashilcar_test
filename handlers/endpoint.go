package handlers

import (
	"context"
	"tashil/models"
)

type EndpointHandler interface {
	RegisterEndpoint(ctx context.Context, endpoint models.RegisterEndpoint) (int, error)
	CallApi(request []byte) error
}

type Notifications interface {
	SendPushNotification(message string, destination string) error
}
