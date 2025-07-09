package database

import (
	"contabi-be/models"
	"database/sql"
)

// ClientsService
type ClientsService struct {
	db *sql.DB
}

// NewClientsService creates a new instance of ClientsService
func NewClientsService(db *sql.DB) *ClientsService {
	return &ClientsService{db: db}
}

// GetClients retrieves all clients
func (cs *ClientsService) GetClients() ([]models.Client, error) {
	q := `
		SELECT 
			c.id
			, c.name
		FROM clients c
		WHERE c.active = true
		ORDER BY name ASC
	`

	rows, err := cs.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var c models.Client
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}

		clients = append(clients, c)
	}

	return clients, nil
}

// GetClientInfo retrieves complete client information using the view
func (cs *ClientsService) GetClientInfo(clientID string) (models.ClientInfo, error) {
	var client models.ClientInfo
	q := `
		SELECT 
			id,
			name,
			rfc,
			clave_ciec,
			clave_fiel,
			fiel_expiration,
			monthly_fee,
			active,
			regimen_id,
			regimen_name,
			supervisor_id,
			supervisor_name,
			responsible_id,
			responsible_name,
			emisor_id,
			emisor_name,
			last_payment_month,
			last_payment_date,
			updated_at
		FROM client_info_view
		WHERE id = $1
	`

	row := cs.db.QueryRow(q, clientID)
	if err := row.Scan(
		&client.ID,
		&client.Name,
		&client.RFC,
		&client.ClaveCIEC,
		&client.ClaveFiel,
		&client.FielExpiration,
		&client.MonthlyFee,
		&client.Active,
		&client.RegimenID,
		&client.RegimenName,
		&client.SupervisorID,
		&client.SupervisorName,
		&client.ResponsibleID,
		&client.ResponsibleName,
		&client.EmisorID,
		&client.EmisorName,
		&client.LastPaymentMonth,
		&client.LastPaymentDate,
		&client.UpdatedAt,
	); err != nil {
		return client, err
	}

	return client, nil
}

// CreateClient creates a new client
func (cs *ClientsService) CreateClient(client models.Client, assignments models.ClientAssignments) error {
	tx, err := cs.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `
		INSERT INTO clients (
			name
			, rfc
			, clave_ciec
			, clave_fiel
			, fiel_expiration
			, monthly_fee
			, regimen_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
		RETURNING id
	`

	row := tx.QueryRow(
		q, client.Name, client.RFC, client.ClaveCIEC, client.ClaveFiel,
		client.FielExpiration, client.MonthlyFee, client.RegimenID,
	)
	if err := row.Scan(&assignments.ClientID); err != nil {
		return err
	}

	queryAssignments := `
		INSERT INTO client_assignments (
			client_id, supervisor_id, responsible_id, emisor_id)
		VALUES ($1, $2, $3, $4)
	`

	_, err = tx.Exec(
		queryAssignments,
		assignments.ClientID,
		assignments.SupervisorID,
		assignments.ResponsibleID,
		assignments.EmisorID,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetAllClientsInfo retrieves all clients with complete information using the view
func (cs *ClientsService) GetAllClientsInfo() ([]models.ClientInfo, error) {
	q := `
		SELECT 
			id,
			name,
			rfc,
			clave_ciec,
			clave_fiel,
			fiel_expiration,
			monthly_fee,
			active,
			regimen_id,
			regimen_name,
			supervisor_id,
			supervisor_name,
			responsible_id,
			responsible_name,
			emisor_id,
			emisor_name,
			last_payment_month,
			last_payment_date,
			updated_at
		FROM client_info_view
		ORDER BY name ASC
	`

	rows, err := cs.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lpm sql.NullString
	var lpd sql.NullString
	var ua sql.NullString
	var clients []models.ClientInfo
	for rows.Next() {
		var client models.ClientInfo
		if err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.RFC,
			&client.ClaveCIEC,
			&client.ClaveFiel,
			&client.FielExpiration,
			&client.MonthlyFee,
			&client.Active,
			&client.RegimenID,
			&client.RegimenName,
			&client.SupervisorID,
			&client.SupervisorName,
			&client.ResponsibleID,
			&client.ResponsibleName,
			&client.EmisorID,
			&client.EmisorName,
			&lpm,
			&lpd,
			&ua,
		); err != nil {
			return nil, err
		}

		if lpm.Valid {
			client.LastPaymentMonth = lpm.String
		}

		if lpd.Valid {
			client.LastPaymentDate = lpd.String
		}
		if ua.Valid {
			client.UpdatedAt = ua.String
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// GetActiveClientsInfo retrieves only active clients with complete information
func (cs *ClientsService) GetActiveClientsInfo() ([]models.ClientInfo, error) {
	q := `
		SELECT 
			id,
			name,
			rfc,
			clave_ciec,
			clave_fiel,
			fiel_expiration,
			monthly_fee,
			active,
			regimen_id,
			regimen_name,
			supervisor_id,
			supervisor_name,
			responsible_id,
			responsible_name,
			emisor_id,
			emisor_name,
			last_payment_month,
			last_payment_date,
			updated_at
		FROM active_clients_view
		ORDER BY name ASC
	`

	rows, err := cs.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lpm sql.NullString
	var lpd sql.NullString
	var ua sql.NullString
	var clients []models.ClientInfo
	for rows.Next() {
		var client models.ClientInfo
		if err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.RFC,
			&client.ClaveCIEC,
			&client.ClaveFiel,
			&client.FielExpiration,
			&client.MonthlyFee,
			&client.Active,
			&client.RegimenID,
			&client.RegimenName,
			&client.SupervisorID,
			&client.SupervisorName,
			&client.ResponsibleID,
			&client.ResponsibleName,
			&client.EmisorID,
			&client.EmisorName,
			&lpm,
			&lpd,
			&ua,
		); err != nil {
			return nil, err
		}

		if lpm.Valid {
			client.LastPaymentMonth = lpm.String
		}

		if lpd.Valid {
			client.LastPaymentDate = lpd.String
		}
		if ua.Valid {
			client.UpdatedAt = ua.String
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// GetClientsWithPendingPayments retrieves clients with pending payments
func (cs *ClientsService) GetClientsWithPendingPayments() ([]models.ClientWithPendingPayment, error) {
	q := `
		SELECT 
			id,
			name,
			rfc,
			regimen_name,
			monthly_fee,
			last_payment_month,
			last_payment_date,
			updated_at,
			supervisor_name,
			responsible_name,
			emisor_name,
			payment_status
		FROM clients_with_pending_payments
		ORDER BY name ASC
	`

	rows, err := cs.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lpm sql.NullString
	var lpd sql.NullString
	var ua sql.NullString
	var clients []models.ClientWithPendingPayment
	for rows.Next() {
		var client models.ClientWithPendingPayment
		if err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.RFC,
			&client.RegimenName,
			&client.MonthlyFee,
			&lpm,
			&lpd,
			&ua,
			&client.SupervisorName,
			&client.ResponsibleName,
			&client.EmisorName,
			&client.PaymentStatus,
		); err != nil {
			return nil, err
		}

		if lpm.Valid {
			client.LastPaymentMonth = lpm.String
		}

		if lpd.Valid {
			client.LastPaymentDate = lpd.String
		}
		if ua.Valid {
			client.UpdatedAt = ua.String
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// UpdateClient updates client information
func (cs *ClientsService) UpdateClient(clientID string, client models.Client) error {
	q := `
		UPDATE clients 
		SET name = $1, rfc = $2, clave_ciec = $3, clave_fiel = $4, 
		    fiel_expiration = $5, monthly_fee = $6, regimen_id = $7, active = $8
		WHERE id = $9
	`

	_, err := cs.db.Exec(q, client.Name, client.RFC, client.ClaveCIEC, client.ClaveFiel,
		client.FielExpiration, client.MonthlyFee, client.RegimenID, client.Active, clientID)

	return err
}

// UpdateClientAssignments updates client assignments (supervisor, responsible, emisor)
func (cs *ClientsService) UpdateClientAssignments(clientID string, assignments models.ClientAssignments) error {
	q := `
		UPDATE client_assignments 
		SET supervisor_id = $1, responsible_id = $2, emisor_id = $3
		WHERE client_id = $4
	`

	_, err := cs.db.Exec(q, assignments.SupervisorID, assignments.ResponsibleID,
		assignments.EmisorID, clientID)

	return err
}

// UpdateClientPayment updates client payment information
func (cs *ClientsService) UpdateClientPayment(clientID string, payment models.ClientPayment) error {
	q := `
		INSERT INTO client_payments (client_id, last_payment_month, last_payment_date)
		VALUES ($1, $2, $3)
		ON CONFLICT (client_id) 
		DO UPDATE SET 
			last_payment_month = EXCLUDED.last_payment_month,
			last_payment_date = EXCLUDED.last_payment_date,
			updated_at = CURRENT_TIMESTAMP
	`

	_, err := cs.db.Exec(q, clientID, payment.LastPaymentMonth, payment.LastPaymentDate)

	return err
}

// DeactivateClient sets a client as inactive (soft delete)
func (cs *ClientsService) DeactivateClient(clientID string) error {
	q := `UPDATE clients SET active = false WHERE id = $1`

	_, err := cs.db.Exec(q, clientID)

	return err
}

// ActivateClient sets a client as active
func (cs *ClientsService) ActivateClient(clientID string) error {
	q := `UPDATE clients SET active = true WHERE id = $1`

	_, err := cs.db.Exec(q, clientID)

	return err
}
