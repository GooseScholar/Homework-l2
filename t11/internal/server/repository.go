package server

import (
	"context"
	"t11/internal/cache"
)

type Repository interface {
	Post(context.Context, string)
	Get(context.Context, string) (*cache.Cache, error)
}
