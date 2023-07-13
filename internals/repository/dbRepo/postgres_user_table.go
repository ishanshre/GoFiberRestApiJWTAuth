package dbrepo

import "github.com/ishanshre/GoFiberRestApiJWTAuth/internals/models"

func (p *postgresDbRepo) AllUsers(limit, offset int) ([]*models.User, error) {
	users := []*models.User{}
	return users, nil
}
