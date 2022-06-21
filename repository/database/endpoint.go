package database

import (
	"context"
	"database/sql"
	"tashil/models"
	"tashil/repository"
)

type endpointRepository struct {
	db *sql.DB
}

func NewEndpointRepository(db *sql.DB) repository.EndpointRepository {
	return &endpointRepository{db: db}
}

func (e *endpointRepository) SaveEndpointRules(ctx context.Context, request models.RegisterEndpoint) error {
	panic("implement me")
}
