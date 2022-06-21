package handlers

import "context"

type Controller interface {
	RegisterRule(ctx context.Context) error
}
