package month

import (
	"context"
	"home-accounting/models"
)

type Repository interface {
	GetById(ctx context.Context, id int32) (models.Month, error)
}
