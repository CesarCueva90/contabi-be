package database

import (
	"contabi-be/models"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// LoginService
type LoginService struct {
	db *sql.DB
}

// NewUsersService creates a new instance of LoginService
func NewLoginService(db *sql.DB) *LoginService {
	return &LoginService{db: db}
}

// Login implements controller.DataBaseService.
func (ls *LoginService) Login(login string, password string) (models.User, error) {
	user := models.User{}

	q := `
		SELECT 
			id
			, username
			, password
			, active
			, ur.role_id
		FROM users u
		INNER JOIN user_role ur ON u.id = ur.user_id
		WHERE u.username = $1
		ORDER BY username ASC LIMIT 1
	`

	rows, err := ls.db.Query(q, login)

	if err != nil {
		return user, err
	}

	defer rows.Close()

	var e models.User
	userFound := false

	for rows.Next() {
		if err := rows.Scan(&e.ID, &e.Username, &e.Password, &e.Active, &e.Role); err != nil {
			return user, err
		}

		user = e
		userFound = true
	}

	// Check if user was found
	if !userFound {
		return models.User{}, fmt.Errorf("user not found")
	}

	// Compare the hashed password with the provided password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password)); err != nil {
		return models.User{}, fmt.Errorf("invalid password")
	}

	return user, nil
}
