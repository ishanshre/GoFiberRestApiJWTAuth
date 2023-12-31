package dbrepo

import (
	"context"
	"fmt"

	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/models"
)

// AllUsers returns slice of all the user from database
func (p *postgresDbRepo) AllUsers(limit, offset int) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	query := `SELECT * FROM users LIMIT $1 OFFSET $2`
	rows, err := p.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("cannot execute the query: %s", err.Error())
	}
	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("cannot scan the row: %s", err.Error())
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByID takes id as parameter and returns the user information matching the id
func (p *postgresDbRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `SELECT * FROM users WHERE id=$1`
	row := p.DB.QueryRowContext(ctx, query, id)
	user := &models.User{}
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByUsername takes username as parameter and returns the user information matching the username
func (p *postgresDbRepo) GetUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `SELECT * FROM users WHERE username=$1`
	row := p.DB.QueryRowContext(ctx, query, username)
	user := &models.User{}
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail takes email as parameter and returns the user information matching the email
func (p *postgresDbRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `SELECT * FROM users WHERE email=$1`
	row := p.DB.QueryRowContext(ctx, query, email)
	user := &models.User{}
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return user, nil
}

// func DeleteUser takes id as parameter and delete the user
func (p *postgresDbRepo) DeleteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	query := `DELETE FROM users WHERE username=$1`
	res, err := p.DB.ExecContext(ctx, query, username)
	if err != nil {
		return fmt.Errorf("error in executing the query: %v", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not deleted or user does not exists")
	}
	return nil
}

// func UpdateUser takes user model as paramter and updates the user
func (p *postgresDbRepo) UpdateUser(u *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	query := `
		UPDATE users
		SET first_name = $2, last_name = $3
		WHERE id=$1
		RETURNING id, first_name, last_name, username, email, password, role, created_at, updated_at
	`
	row := p.DB.QueryRowContext(
		ctx,
		query,
		u.ID,
		u.FirstName,
		u.LastName,
	)
	user := &models.User{}
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateRole takes id and role as parameter and update user roles
func (p *postgresDbRepo) UpdateRole(u *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	query := `
		UPDATE users
		SET role=$2
		WHERE id=$1
		RETURNING id, first_name, last_name, username, email, password, role, created_at, updated_at
	`
	row := p.DB.QueryRowContext(
		ctx,
		query,
		u.ID,
		u.Role,
	)
	user := &models.User{}
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser takes user model as paramter to insert user into datbase
func (p *postgresDbRepo) CreateUser(user *models.ValidateUser) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	query := `
		INSERT INTO users (first_name, last_name, username, email, password, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		RETURNING id, first_name, last_name, username, email, password, role, created_at, updated_at
	`
	row := p.DB.QueryRowContext(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	u := &models.User{}
	if err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (p *postgresDbRepo) UsernameExists(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `
		SELECT COUNT(*) FROM users WHERE username=$1
	`
	var count int
	if err := p.DB.QueryRowContext(ctx, query, username).Scan(&count); err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	return count > 0, nil
}

func (p *postgresDbRepo) EmailExists(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `
		SELECT COUNT(*) FROM users WHERE email=$1
	`
	var count int
	if err := p.DB.QueryRowContext(ctx, query, email).Scan(&count); err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	return count > 0, nil
}
