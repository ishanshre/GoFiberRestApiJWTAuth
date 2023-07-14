package dbrepo

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/models"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) *models.User {
	args := &models.User{
		FirstName: utils.RandomString(6),
		LastName:  utils.RandomString(6),
		Username:  utils.RandomString(6),
		Email:     utils.RandomString(6),
		Password:  utils.RandomString(6),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err := testQueries.CreateUser(args)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.NotZero(t, user.Role)
	require.Equal(t, args.FirstName, user.FirstName)
	require.Equal(t, args.LastName, user.LastName)
	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.Password, user.Password)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestCreateUser_Failure(t *testing.T) {
	args := &models.User{
		FirstName: "test",
		LastName:  "test",
		Username:  "test123",
		Email:     "test1234",
		Password:  "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	testQueries.CreateUser(args)
	_, err := testQueries.CreateUser(args)
	require.Error(t, err, `pq: duplicate key value violates unique constraint "users_username_key"`)
}

func TestGetUserByID(t *testing.T) {
	createUser := createRandomUser(t)
	getUser, err := testQueries.GetUserByID(createUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getUser)
	require.NotEmpty(t, getUser.Role)
	require.Equal(t, createUser.ID, getUser.ID)
	require.Equal(t, createUser.FirstName, getUser.FirstName)
	require.Equal(t, createUser.LastName, getUser.LastName)
	require.Equal(t, createUser.Username, getUser.Username)
	require.Equal(t, createUser.Email, getUser.Email)
	require.Equal(t, createUser.Password, getUser.Password)
}

func TestGetUserByID_Failure(t *testing.T) {
	_, err := testQueries.GetUserByID(int(utils.RandomInt(10000000000, 10000000000)))
	require.Error(t, err)
}
func TestGetUserByUsername(t *testing.T) {
	createUser := createRandomUser(t)
	getUser, err := testQueries.GetUserByUsername(createUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, getUser)
	require.NotEmpty(t, getUser.Role)
	require.Equal(t, createUser.ID, getUser.ID)
	require.Equal(t, createUser.FirstName, getUser.FirstName)
	require.Equal(t, createUser.LastName, getUser.LastName)
	require.Equal(t, createUser.Username, getUser.Username)
	require.Equal(t, createUser.Email, getUser.Email)
	require.Equal(t, createUser.Password, getUser.Password)
}

func TestGetUserByUsername_Failure(t *testing.T) {
	_, err := testQueries.GetUserByUsername(utils.RandomString(10))
	require.Error(t, err)
}
func TestGetUserByEmail(t *testing.T) {
	createUser := createRandomUser(t)
	getUser, err := testQueries.GetUserByEmail(createUser.Email)
	require.NoError(t, err)
	require.NotEmpty(t, getUser)
	require.NotEmpty(t, getUser.Role)
	require.Equal(t, createUser.ID, getUser.ID)
	require.Equal(t, createUser.FirstName, getUser.FirstName)
	require.Equal(t, createUser.LastName, getUser.LastName)
	require.Equal(t, createUser.Username, getUser.Username)
	require.Equal(t, createUser.Email, getUser.Email)
	require.Equal(t, createUser.Password, getUser.Password)
}

func TestGetUserByEmail_Failure(t *testing.T) {
	_, err := testQueries.GetUserByEmail(utils.RandomString(10))
	require.Error(t, err)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testQueries.DeleteUser(user.Username)
	require.NoError(t, err)
	getUser, err := testQueries.GetUserByUsername(user.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getUser)
}

func TestAllUsers(t *testing.T) {
	for i := 0; i < 20; i++ {
		createRandomUser(t)
	}
	limit := 20
	offset := 0
	users, err := testQueries.AllUsers(limit, offset)
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Len(t, users, limit)
}

func TestAllUsers_Failure(t *testing.T) {
	for i := 0; i < 20; i++ {
		createRandomUser(t)
	}
	limit := -1
	offset := 100
	users, err := testQueries.AllUsers(limit, offset)
	assert.Error(t, err)
	assert.True(t, true, fmt.Errorf("cannot execute the query: %s", err.Error()))
	assert.Nil(t, users)

	query := `
		INSERT INTO users (username, email, password, created_at, updated_at)
		VALUES ($2,$3,'srehesdfg', $1, $1)
	`
	_, err = testQueries.DB.Exec(query, time.DateOnly, utils.RandomString(7), utils.RandomString(10))
	assert.NoError(t, err)
	limit = 10000000
	offset = 0
	_, err = testQueries.AllUsers(limit, offset)
	assert.Error(t, err)
}

func TestUpdateUser(t *testing.T) {
	createUser := createRandomUser(t)
	updateUser := &models.User{
		ID:        createUser.ID,
		FirstName: utils.RandomString(12),
		LastName:  utils.RandomString(10),
	}
	updated, err := testQueries.UpdateUser(updateUser)
	require.NoError(t, err)
	require.NotEmpty(t, updated)
	require.Equal(t, createUser.ID, updated.ID)
	require.Equal(t, createUser.Username, updated.Username)
	require.Equal(t, createUser.Email, updated.Email)
	require.Equal(t, createUser.Role, updated.Role)
	require.NotEqual(t, createUser.FirstName, updated.FirstName)
	require.NotEqual(t, createUser.LastName, updated.LastName)
}

func TestUpdateUser_Failure(t *testing.T) {
	query := `
		INSERT INTO users (id, username, email, password, created_at, updated_at)
		VALUES ($4, $2,$3,'srehesdfg', $1, $1)
	`
	id := int(utils.RandomInt(10000000000, 10000000000))
	_, _ = testQueries.DB.Exec(query, time.DateOnly, utils.RandomString(7), utils.RandomString(10), id)
	updateUser := &models.User{
		ID:        id,
		FirstName: "",
	}
	updated, err := testQueries.UpdateUser(updateUser)
	assert.Error(t, err)
	assert.Nil(t, updated)

}
func TestUpdateRole(t *testing.T) {
	createUser := createRandomUser(t)
	updateUser := &models.User{
		ID:   createUser.ID,
		Role: 3,
	}
	updated, err := testQueries.UpdateRole(updateUser)
	require.NoError(t, err)
	require.NotEmpty(t, updated)
	require.Equal(t, createUser.ID, updated.ID)
	require.Equal(t, createUser.Username, updated.Username)
	require.Equal(t, createUser.Email, updated.Email)
	require.Equal(t, createUser.FirstName, updated.FirstName)
	require.NotEqual(t, createUser.Role, updated.Role)
}
func TestUpdateRole_Failure(t *testing.T) {
	query := `
		INSERT INTO users (id, username, email, password, created_at, updated_at)
		VALUES ($4, $2,$3,'srehesdfg', $1, $1)
	`
	id := int(utils.RandomInt(10000000000, 10000000000))
	_, _ = testQueries.DB.Exec(query, time.DateOnly, utils.RandomString(7), utils.RandomString(10), id)
	updateUser := &models.User{
		ID:   id,
		Role: 3,
	}
	updated, err := testQueries.UpdateRole(updateUser)
	assert.Error(t, err)
	assert.Nil(t, updated)

}
