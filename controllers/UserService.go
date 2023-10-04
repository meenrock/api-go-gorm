package controllers

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserFoodRepository struct {
	DB *sql.DB
}

func NewUserFoodRepository(db *sql.DB) *UserFoodRepository {
	return &UserFoodRepository{DB: db}
}

func (r *UserFoodRepository) GetUserFavFoodByUserId(id uuid.UUID) (*sql.Rows, error) {
	rows, err := r.DB.Query("SELECT "+
		"f.id, "+
		"f.name, "+
		"f.energy, "+
		"f.protein, "+
		"f.carbohydrate, "+
		"f.fat, "+
		"f.sodium, "+
		"f.cholesterol, "+
		"f.created_at, "+
		"f.updated_at "+
		"FROM food uf "+
		"inner join food f on uf.food_id = f.id "+
		"WHERE uf.user_id = $1 and uf.user_food_type = 'LIKE'", id)

	if err != nil {
		return nil, err
	}

	return rows, nil
}
