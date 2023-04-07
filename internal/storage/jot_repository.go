package storage

import (
	"context"
	"fmt"

	"github.com/salty-max/grest/internal/models"
)

type JotRepository struct {
	db *Database
}

func NewJotRepository(db *Database) *JotRepository {
	return &JotRepository{
		db: db,
	}
}

func (j *JotRepository) GetJots(ctx context.Context) ([]models.Jot, error) {
	var jots []models.Jot
	result := j.db.Client.Find(&jots)
	if result.Error != nil {
		return nil, fmt.Errorf("error fetching jots: %w", result.Error)
	}

	return jots, nil
}

func (j *JotRepository) CreateJot(ctx context.Context, jot models.Jot) (models.Jot, error) {
	if err := j.db.Client.Create(&jot).Error; err != nil {
		return models.Jot{}, fmt.Errorf("error creating jot: %w", err)
	}

	return jot, nil
}
