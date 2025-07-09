package database

import (
	"contabi-be/models"
	"database/sql"
	"fmt"
)

// MenusService
type MenusService struct {
	db *sql.DB
}

// NewMenusService creates a new instance of MenusService
func NewMenusService(db *sql.DB) *MenusService {
	return &MenusService{db: db}
}

// GetEmisors retrieves all emisors
func (ms *MenusService) GetEmisors() ([]models.Emisor, error) {
	q := `
		SELECT 
			e.id
			, e.name
		FROM emisors e
		WHERE e.active = true
		ORDER BY name ASC
	`

	rows, err := ms.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emisors []models.Emisor
	for rows.Next() {
		var e models.Emisor
		if err := rows.Scan(&e.ID, &e.Name); err != nil {
			return nil, err
		}

		emisors = append(emisors, e)
	}

	fmt.Println(emisors)

	return emisors, nil
}

// GetSupervisors retrieves all supervisors
func (ms *MenusService) GetSupervisors() ([]models.Supervisor, error) {
	q := `
		SELECT 
			u.id
			, u.username
		FROM users u
		INNER JOIN user_roles ur ON ur.user_id = u.id
		WHERE u.active = true
		AND ur.role_id = 2
		ORDER BY username ASC
	`

	rows, err := ms.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var supervisors []models.Supervisor
	for rows.Next() {
		var s models.Supervisor
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}

		supervisors = append(supervisors, s)
	}

	return supervisors, nil
}

// GetResponsiblesBySupervisor retrieves all responsibles
func (ms *MenusService) GetResponsiblesBySupervisor(supervisorID string) ([]models.Responsible, error) {
	q := `
		SELECT 
			u.id
			, u.username
		FROM users u
		INNER JOIN supervisor_responsibles sr ON sr.responsible_id = u.id
		WHERE u.active = true
		AND sr.supervisor_id = $1
		ORDER BY username ASC
	`

	rows, err := ms.db.Query(q, supervisorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responsibles []models.Responsible
	for rows.Next() {
		var r models.Responsible
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			return nil, err
		}

		responsibles = append(responsibles, r)
	}

	return responsibles, nil
}

// GetRegimenes retrieves all regimenes
func (ms *MenusService) GetRegimenes() ([]models.Regimen, error) {
	q := `
		SELECT 
			r.id
			, r.name
		FROM regimenes r
		WHERE r.active = true
		ORDER BY name ASC
	`

	rows, err := ms.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regimenes []models.Regimen
	for rows.Next() {
		var r models.Regimen
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			return nil, err
		}

		regimenes = append(regimenes, r)
	}

	return regimenes, nil
}
