package database

import (
	"contabi-be/models"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// AccountancyService
type AccountancyService struct {
	db *sql.DB
}

// NewAccountancyService creates a new instance of AccountancyService
func NewAccountancyService(db *sql.DB) *AccountancyService {
	return &AccountancyService{db: db}
}

// GetClientsBySupervisor retrieves all the clients of a specific supervisor
func (as *AccountancyService) GetClientsBySupervisor(supervisorID string) ([]models.AccountancyClientInfo, error) {
	q := `
		SELECT DISTINCT
			id,
			name,
			rfc,
			clave_ciec,
			clave_fiel,
			fiel_expiration,
			regimen_id,
			regimen_name,
			responsible_id,
			responsible_name,
			emisor_id,
			emisor_name
		FROM client_info_view
		WHERE supervisor_id = $1
		AND active = true
	`
	rows, err := as.db.Query(q, supervisorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.AccountancyClientInfo
	for rows.Next() {
		var client models.AccountancyClientInfo
		if err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.RFC,
			&client.ClaveCIEC,
			&client.ClaveFiel,
			&client.FielExpiration,
			&client.RegimenID,
			&client.RegimenName,
			&client.ResponsibleID,
			&client.ResponsibleName,
			&client.EmisorID,
			&client.EmisorName,
		); err != nil {
			return clients, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// GetClientAssignmentsMatrix retrieves, for all active clients, the list of assignment types and whether each client has each assignment
func (as *AccountancyService) GetClientAssignmentsMatrix() ([]models.ClientAssignmentMatrixRow, error) {
	q := `
		SELECT
			c.id AS client_id,
			c.name AS client_name,
			at.id AS assignment_type_id,
			at.name AS assignment_type_name,
			CASE WHEN cat.assignment_type_id IS NULL THEN FALSE ELSE TRUE END AS selected
		FROM clients c
		CROSS JOIN accountancy_types at
		LEFT JOIN client_assignments_types cat
		  ON c.id = cat.client_id AND at.id = cat.assignment_type_id
		WHERE c.active = TRUE
		ORDER BY c.name, at.name;
	`
	rows, err := as.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.ClientAssignmentMatrixRow
	for rows.Next() {
		var row models.ClientAssignmentMatrixRow
		if err := rows.Scan(
			&row.ClientID,
			&row.ClientName,
			&row.AssignmentTypeID,
			&row.AssignmentTypeName,
			&row.Selected,
		); err != nil {
			return result, err
		}
		result = append(result, row)
	}
	return result, nil
}

// UpdateClientAssignments updates assignments of a client according to the state
func (as *AccountancyService) UpdateClientAssignments(clientID string, assignments []models.AssignmentSelection) error {
	var toDelete []int
	var toAdd []int
	for _, a := range assignments {
		if a.Selected {
			toAdd = append(toAdd, a.AssignmentTypeID)
		} else {
			toDelete = append(toDelete, a.AssignmentTypeID)
		}
	}

	if len(toDelete) > 0 {
		_, err := as.db.Exec(`
			DELETE FROM client_assignments_types
			WHERE client_id = $1 AND assignment_type_id = ANY($2)
		`, clientID, pq.Array(toDelete))
		if err != nil {
			return err
		}
	}

	for _, assignmentTypeID := range toAdd {
		_, err := as.db.Exec(`
			INSERT INTO client_assignments_types (client_id, assignment_type_id)
			VALUES ($1, $2)
			ON CONFLICT DO NOTHING
		`, clientID, assignmentTypeID)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetClientsByResonsible retrieves all the clients of a specific responsible
func (as *AccountancyService) GetClientsByResonsible(responsibleID string) ([]models.AccountancyClientInfo, error) {
	q := `
		SELECT DISTINCT
			id,
			name,
			rfc,
			clave_ciec,
			clave_fiel,
			fiel_expiration,
			regimen_id,
			regimen_name,
			responsible_id,
			responsible_name,
			emisor_id,
			emisor_name
		FROM client_info_view
		WHERE responsible_id = $1
		AND active = true
	`
	rows, err := as.db.Query(q, responsibleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.AccountancyClientInfo
	for rows.Next() {
		var client models.AccountancyClientInfo
		if err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.RFC,
			&client.ClaveCIEC,
			&client.ClaveFiel,
			&client.FielExpiration,
			&client.RegimenID,
			&client.RegimenName,
			&client.ResponsibleID,
			&client.ResponsibleName,
			&client.EmisorID,
			&client.EmisorName,
		); err != nil {
			return clients, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// CreateClientAccountancyStatusWithAssignments creates a new monthly record for a client
func (as *AccountancyService) CreateClientAccountancyStatusWithAssignments(status models.ClientAccountancyStatus, assignments []models.ClientAccountancyAssignment) error {
	tx, err := as.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var statusID int
	err = tx.QueryRow(
		`INSERT INTO client_accountancy_status (client_id, month, due_date, observaciones)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		status.ClientID, status.Month, status.DueDate, status.Observacion,
	).Scan(&statusID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, a := range assignments {
		_, err := tx.Exec(
			`INSERT INTO client_accountancy_assignments (status_id, assignment_type_id, assignment_status_id)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (status_id, assignment_type_id) DO UPDATE SET assignment_status_id = EXCLUDED.assignment_status_id`,
			statusID, a.AssignmentTypeID, a.AssignmentStatusID,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// UpdateClientAccountancyStatusWithAssignments updates an existing monthly record for a client
func (as *AccountancyService) UpdateClientAccountancyStatusWithAssignments(statusID int, clientID string, status models.ClientAccountancyStatus, assignments []models.ClientAccountancyAssignment) error {
	tx, err := as.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Update main status, validating that it belongs to the client
	result, err := tx.Exec(
		`UPDATE client_accountancy_status 
		 SET due_date = $1, observaciones = $2
		 WHERE id = $3 AND client_id = $4`,
		status.DueDate, status.Observacion, statusID, clientID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Verify that at least one row was affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("status not found or does not belong to the specified client")
	}

	// Actualizar las asignaciones
	for _, a := range assignments {
		_, err := tx.Exec(
			`INSERT INTO client_accountancy_assignments (status_id, assignment_type_id, assignment_status_id)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (status_id, assignment_type_id) 
			 DO UPDATE SET assignment_status_id = EXCLUDED.assignment_status_id`,
			statusID, a.AssignmentTypeID, a.AssignmentStatusID,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// GetClientAccountancyHistory gets the hisotory for a client accountancy behavior
func (as *AccountancyService) GetClientAccountancyHistory(clientID string) (models.ClientAccountancyHistoryWithAssignments, error) {
	var result models.ClientAccountancyHistoryWithAssignments

	qActiveAssignments := `
		SELECT at.id, at.name
		FROM client_assignments_types cat
		JOIN accountancy_types at ON cat.assignment_type_id = at.id
		WHERE cat.client_id = $1
		ORDER BY at.name ASC
	`
	rowsAA, err := as.db.Query(qActiveAssignments, clientID)
	if err != nil {
		return result, err
	}
	defer rowsAA.Close()
	for rowsAA.Next() {
		var at models.AccountancyType
		if err := rowsAA.Scan(&at.ID, &at.Name); err != nil {
			return result, err
		}
		result.ActiveAssignments = append(result.ActiveAssignments, at)
	}

	qStatus := `
		SELECT 
			id
			, client_id
			, month
			, due_date
			, observaciones 
		FROM client_accountancy_status 
		WHERE client_id = $1 
		ORDER BY month DESC
	`
	rows, err := as.db.Query(qStatus, clientID)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var status models.ClientAccountancyStatus
		if err := rows.Scan(&status.ID, &status.ClientID, &status.Month, &status.DueDate, &status.Observacion); err != nil {
			return result, err
		}

		// Gets assignmentsfor this status, only the actives for the client
		qAssign := `
			SELECT 
				caa.id
				, caa.status_id
				, caa.assignment_type_id
				, at.name AS assignment_type_name
				, caa.assignment_status_id
				, ast.name AS assignment_status_name
			FROM client_accountancy_assignments caa
			JOIN accountancy_types at ON caa.assignment_type_id = at.id
			JOIN assignment_statuses ast ON caa.assignment_status_id = ast.id
			JOIN client_assignments_types cat ON caa.assignment_type_id = cat.assignment_type_id AND cat.client_id = $2
			WHERE caa.status_id = $1
		`

		assignRows, err := as.db.Query(qAssign, status.ID, status.ClientID)
		if err != nil {
			return result, err
		}
		var assignments []models.ClientAccountancyAssignment
		for assignRows.Next() {
			var a models.ClientAccountancyAssignment
			if err := assignRows.Scan(&a.ID, &a.StatusID, &a.AssignmentTypeID, &a.AssignmentTypeName, &a.AssignmentStatusID, &a.AssignmentStatusName); err != nil {
				assignRows.Close()
				return result, err
			}
			assignments = append(assignments, a)
		}
		assignRows.Close()

		result.History = append(result.History, models.ClientAccountancyHistoryEntry{
			Status:      status,
			Assignments: assignments,
		})
	}

	return result, nil
}
