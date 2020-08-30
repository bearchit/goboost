package crud

import (
	"context"
)

type Creator interface {
	Create(ctx context.Context, e interface{}) error
}

type Updater interface {
	Update(ctx context.Context, e interface{}) error
}

type Deleter interface {
	Delete(ctx context.Context, e interface{}) error
}

type Cond map[string]interface{}

type Finder interface {
	FindByPK(ctx context.Context, pk interface{}, result interface{}) error
	FindBy(ctx context.Context, cond Cond, result interface{}) error
}

type Repository interface {
	Creator
	Updater
	Deleter
	Finder
}
