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

// NewLoginService creates a new instance of LoginService
func NewLoginService(db *sql.DB) *LoginService {
	return &LoginService{db: db}
}

// Login implements controller.DataBaseService.
func (ls *LoginService) Login(login string, password string) (models.User, error) {
	user := models.User{}

	q := `
		SELECT 
			u.id
			, u.username
			, u.password
			, u.active
			, ur.role_id
		FROM users u
		INNER JOIN user_roles ur ON u.id = ur.user_id
		WHERE u.username = $1 AND u.active = true
		LIMIT 1
	`

	row := ls.db.QueryRow(q, login)

	var roleID int
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Active, &roleID); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found")
		}
		return models.User{}, err
	}

	user.Role = roleID

	// Compare the hashed password with the provided password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Printf("DEBUG: username=%s, password='%s', hash='%s', hashlen=%d, passlen=%d\n", user.Username, password, user.Password, len(user.Password), len(password))
		return models.User{}, fmt.Errorf("invalid password")
	}

	return user, nil
}
