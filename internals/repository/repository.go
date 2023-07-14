package repository

import "github.com/ishanshre/GoFiberRestApiJWTAuth/internals/models"

type DatabaseRepo interface {
	AllUsers(limit, offset int) ([]*models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(username string) error
	UpdateUser(u *models.User) (*models.User, error)
	UpdateRole(u *models.User) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UsernameExists(username string) (bool, error)
	EmailExists(email string) (bool, error)
}
