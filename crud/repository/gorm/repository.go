package gorm

import (
	"context"

	"github.com/bearchit/goboost/crud"
	"github.com/bearchit/goboost/structs"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) crud.Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Create(ctx context.Context, e interface{}) error {
	return r.db.Create(e).Error
}

func (r Repository) Update(ctx context.Context, e interface{}) error {
	return r.db.Save(e).Error
}

func (r Repository) Delete(ctx context.Context, e interface{}) error {
	return r.db.Delete(e).Error
}

func (r Repository) FindByPK(ctx context.Context, pk interface{}, result interface{}) error {
	pkField, err := structs.FieldByTagKeyNamePairs(result,
		[]string{"crud", "pk", "gorm", "primary_key"})
	if err != nil {
		return err
	}

	cond := make(map[string]interface{})
	cond[pkField.Name] = pkField.Value
	return r.db.Where(cond).First(result).Error
}

func (r Repository) FindBy(ctx context.Context, cond crud.Cond, result interface{}) error {
	return r.db.Where(cond).First(&result).Error
}
