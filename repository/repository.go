package repository

import (
	"context"
)

type Repository struct {
	ctx    context.Context
	logger string
	client int
}
