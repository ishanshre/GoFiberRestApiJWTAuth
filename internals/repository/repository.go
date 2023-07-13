package repository

import "github.com/ishanshre/GoFiberRestApiJWTAuth/internals/models"

type DatabaseRepo interface {
	AllUsers(limit, offset int) ([]*models.User, error)
}
