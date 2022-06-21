package repository

import (
	"context"
	"tashil/models"
)

type EndpointRepository interface {
	SaveEndpointRules(ctx context.Context, request models.RegisterEndpoint) error
}
