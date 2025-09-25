package database

import (
	"contabi-be/models"
	"database/sql"
)

// NominasService
type NominasService struct {
	db *sql.DB
}

// NewNominasService creates a new instance of NominasService
func NewNominasService(db *sql.DB) *NominasService {
	return &NominasService{db: db}
}

// CreateClientPaymentRecord creates a new client payment record
func (ns *NominasService) CreateClientPaymentRecord(clientPaymentRecord models.ClientHRPayment) error {
	q := `
		INSERT INTO client_hr_payments (
			client_id
			, hr_entity_id
			, payment_month
			, amount
			, paid
			, month
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err := ns.db.Exec(
		q, clientPaymentRecord.ClientID, clientPaymentRecord.HREntityID,
		clientPaymentRecord.PaymentMonth, clientPaymentRecord.Amount,
		clientPaymentRecord.Paid, clientPaymentRecord.Month,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetClientsWithPendingPaymentsByHREntityID gets the list of all clients with pending payments of specific HR entity
func (ns *NominasService) GetClientsWithPendingPaymentsByHREntityID(hrEntityID string) ([]models.ClientWithPendingHRPayment, error) {
	q := `
		SELECT DISTINCT
			c.id AS client_id,
			c.name AS client_name,
			chp.hr_entity_id
		FROM client_hr_payments chp
		JOIN clients c ON chp.client_id = c.id
		WHERE chp.hr_entity_id = $1
		AND chp.paid = FALSE
		ORDER BY c.name ASC;
	`

	rows, err := ns.db.Query(q, hrEntityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.ClientWithPendingHRPayment
	for rows.Next() {
		var c models.ClientWithPendingHRPayment
		if err := rows.Scan(&c.ClientID, &c.ClientName, &c.HREntityID); err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}

	return clients, nil
}

// GetClientPendingPaymentsByHREntityIDDetails gets all the pending payments of a specific client and HR entity
func (ns *NominasService) GetClientPendingPaymentsByHREntityIDDetails(clientID, hrEntityID string) ([]models.ClientWithPendingHRPaymentDetails, error) {
	q := `
		SELECT
			chp.id AS payment_id,
			c.id AS client_id,
			c.name AS client_name,
			chp.hr_entity_id,
			chp.payment_month,
			chp.amount,
			chp.paid,
			chp.month
		FROM client_hr_payments chp
		JOIN clients c ON chp.client_id = c.id
		WHERE chp.hr_entity_id = $1
		AND c.id = $2
		AND chp.paid = FALSE
		ORDER BY chp.payment_month DESC, c.name ASC;
	`

	rows, err := ns.db.Query(q, hrEntityID, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.ClientWithPendingHRPaymentDetails
	for rows.Next() {
		var p models.ClientWithPendingHRPaymentDetails
		if err := rows.Scan(
			&p.ID, &p.ClientID, &p.ClientName, &p.HREntityID,
			&p.PaymentMonth, &p.Amount, &p.Paid, &p.Month,
		); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}

	return payments, nil
}

// UpdateClientPaymentRecord updates a client payment record
func (ns *NominasService) UpdateClientPaymentRecord(clientPaymentRecord models.UpdateClientHRPayment) error {
	q := `
		UPDATE client_hr_payments 
		SET paid = $1, payment_month = $2, amount = $3, month = $4
		WHERE id = $5
	`

	_, err := ns.db.Exec(
		q, clientPaymentRecord.Paid, clientPaymentRecord.PaymentMonth,
		clientPaymentRecord.Amount, clientPaymentRecord.Month, clientPaymentRecord.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetClientHRPaymentsHistory gets all the payments of a specific client and HR entity
func (ns *NominasService) GetClientHRPaymentsHistory(clientID, hrEntityID string) ([]models.ClientHRPayment, error) {
	q := `
		SELECT
			chp.id,
			chp.payment_month,
			chp.amount,
			chp.paid,
			chp.month
		FROM client_hr_payments chp
		WHERE chp.hr_entity_id = $1
		AND chp.client_id = $2
		ORDER BY chp.payment_month DESC;
	`

	rows, err := ns.db.Query(q, hrEntityID, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.ClientHRPayment
	for rows.Next() {
		var p models.ClientHRPayment
		if err := rows.Scan(
			&p.ID, &p.PaymentMonth, &p.Amount, &p.Paid, &p.Month,
		); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}

	return payments, nil
}
