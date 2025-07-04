package database

import (
	"database/sql"
	"fmt"

	"contabi-be/models"
)

// UsersService
type UsersService struct {
	db *sql.DB
}

// NewUsersService creates a new instance of UsersService
func NewUsersService(db *sql.DB) *UsersService {
	return &UsersService{db: db}
}

// GetUsers retrieves all users
func (us *UsersService) GetUsers() ([]models.User, error) {
	q := `
		SELECT 
			u.id
			, u.username
			, u.branch_id
			, u.active
			, ur.role_id
			, b.name
		FROM users u
		INNER JOIN user_role ur ON u.id = ur.user_id
		LEFT JOIN branches b ON u.branch_id = b.id
		ORDER BY username ASC
	`

	rows, err := us.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Active, &u.Role); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

// GetUserInfo retrieves user info by ID
func (us *UsersService) GetUserByID(userID string) (models.User, error) {
	q := `
		SELECT 
			u.id
			, u.username
			, u.active
			, ur.role_id
		FROM users u
		INNER JOIN user_role ur ON u.id = ur.user_id
		WHERE u.id = $1
		ORDER BY u.username ASC LIMIT 1
	`

	row := us.db.QueryRow(q, userID)
	var u models.User
	if err := row.Scan(&u.ID, &u.Username, &u.Active, &u.Role); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return u, err
	}

	return u, nil
}

// CreateUser creates a new user
func (us *UsersService) CreateUser(user models.User) error {
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `
		INSERT INTO users (username, password, active)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	row := tx.QueryRow(q, user.Username, user.Password, user.Active)
	var id string
	if err := row.Scan(&id); err != nil {
		return err
	}

	queryRole := `
		INSERT INTO user_role (user_id, role_id)
		VALUES ($1, $2)
	`

	_, err = tx.Exec(queryRole, id, user.Role)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// UpdateUser updates an user
func (us *UsersService) UpdateUser(user models.User) error {
	q := `
		UPDATE users 
		SET username = $1
		WHERE id = $2
	`

	_, err := us.db.Exec(q, user.Username, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserRole updates the user role
func (us *UsersService) UpdateUserRole(user models.User) error {
	q := `
		UPDATE user_role
		SET  role_id= $1
		WHERE user_id = $2
	`

	_, err := us.db.Exec(q, user.Role, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// PutUserPassword updtaes the user password
func (us *UsersService) PutUserPassword(user models.User) error {
	query := `
		UPDATE users 
		SET password = $1
		WHERE id = $2
		`

	_, err := us.db.Exec(query, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes an user
func (us *UsersService) DeleteUser(userID string) error {
	q := `
		UPDATE users 
		SET active = false 
		WHERE id = $1
	`
	fmt.Println(userID)
	_, err := us.db.Exec(q, userID)
	if err != nil {
		return err
	}

	return nil
}

// GetRoles gets all the roles
func (us *UsersService) GetRoles() ([]models.Role, error) {
	q := `
		SELECT 
			id
			, name
		FROM roles
		ORDER BY name ASC
	`

	rows, err := us.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var r models.Role
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}

	return roles, nil
}
