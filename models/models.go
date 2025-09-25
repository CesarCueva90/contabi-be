package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Active   bool   `json:"active"`
	Role     int    `json:"role"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Emisor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Supervisor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Regimen struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Responsible struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	SupervisorID string `json:"supervisor_id"`
}

type Client struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	RegimenID      string `json:"regimen_id"`
	RFC            string `json:"rfc"`
	ClaveCIEC      string `json:"clave_ciec"`
	ClaveFiel      string `json:"clave_fiel"`
	FielExpiration string `json:"fiel_expiration"`
	MonthlyFee     string `json:"monthly_fee"`
	Active         string `json:"active"`
}

type ClientAssignments struct {
	ClientID        string `json:"client_id"`
	SupervisorID    string `json:"supervisor_id"`
	SupervisorName  string `json:"supervisor_name"`
	ResponsibleID   string `json:"responsible_id"`
	ResponsibleName string `json:"responsible_name"`
	EmisorID        string `json:"emisor_id"`
	EmisorName      string `json:"emisor_name"`
}

type ClientPayment struct {
	ClientID         string `json:"client_id"`
	ClientName       string `json:"client_name"`
	LastPaymentMonth string `json:"last_payment_month"`
	LastPaymentDate  string `json:"last_payment_date"`
	FolioFactura     string `json:"folio_factura"`
	UpdatedAt        string `json:"updated_at"`
}

// ClientInfo represents the full clients info from view
type ClientInfo struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	RFC              string `json:"rfc"`
	RegimenName      string `json:"regimen_name"`
	RegimenID        string `json:"regimen_id"`
	ClaveCIEC        string `json:"clave_ciec"`
	ClaveFiel        string `json:"clave_fiel"`
	FielExpiration   string `json:"fiel_expiration"`
	MonthlyFee       string `json:"monthly_fee"`
	Active           bool   `json:"active"`
	SupervisorID     string `json:"supervisor_id,omitempty"`
	SupervisorName   string `json:"supervisor_name,omitempty"`
	ResponsibleID    string `json:"responsible_id,omitempty"`
	ResponsibleName  string `json:"responsible_name,omitempty"`
	EmisorID         string `json:"emisor_id,omitempty"`
	EmisorName       string `json:"emisor_name,omitempty"`
	LastPaymentMonth string `json:"last_payment_month,omitempty"`
	LastPaymentDate  string `json:"last_payment_date,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}

// ClientWithPendingPayment used for clients with pending payments
type ClientWithPendingPayment struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	RFC              string `json:"rfc"`
	RegimenName      string `json:"regimen_name"`
	MonthlyFee       string `json:"monthly_fee"`
	LastPaymentMonth string `json:"last_payment_month,omitempty"`
	LastPaymentDate  string `json:"last_payment_date,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	SupervisorName   string `json:"supervisor_name,omitempty"`
	ResponsibleName  string `json:"responsible_name,omitempty"`
	EmisorName       string `json:"emisor_name,omitempty"`
	PaymentStatus    string `json:"payment_status"`
}

type ClientPaymentHistory struct {
	ClientID         string `json:"client_id"`
	ClientName       string `json:"client_name"`
	LastPaymentMonth string `json:"last_payment_month"`
	LastPaymentDate  string `json:"last_payment_date"`
	MonthlyFee       string `json:"monthly_fee"`
	FolioFactura     string `json:"folio_factura,omitempty"`
}

type ClientHRPayment struct {
	ID           string  `json:"id"`
	ClientID     string  `json:"client_id"`
	HREntityID   string  `json:"hr_entity_id"`
	PaymentMonth string  `json:"payment_month"`
	Amount       float64 `json:"amount"`
	Paid         bool    `json:"paid"`
	Month        string  `json:"month"`
}

type ClientWithPendingHRPaymentDetails struct {
	ID           string  `json:"payment_id"`
	ClientID     string  `json:"client_id"`
	ClientName   string  `json:"client_name"`
	HREntityID   string  `json:"hr_entity_id"`
	PaymentMonth string  `json:"payment_month"`
	Amount       float64 `json:"amount"`
	Paid         bool    `json:"paid"`
	Month        string  `json:"month"`
}

type ClientWithPendingHRPayment struct {
	ClientID   string `json:"client_id"`
	ClientName string `json:"client_name"`
	HREntityID string `json:"hr_entity_id"`
}

type UpdateClientHRPayment struct {
	ID   string `json:"id"`
	Paid string `json:"paid"`
}

// ACCOUNTANCY

type AccountancyType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccountancyAssignmentStatus struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AccountancyClientInfo represents the accountancy info of all the clients
type AccountancyClientInfo struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	RFC             string `json:"rfc"`
	RegimenName     string `json:"regimen_name"`
	RegimenID       string `json:"regimen_id"`
	ClaveCIEC       string `json:"clave_ciec"`
	ClaveFiel       string `json:"clave_fiel"`
	FielExpiration  string `json:"fiel_expiration"`
	SupervisorID    string `json:"supervisor_id,omitempty"`
	SupervisorName  string `json:"supervisor_name,omitempty"`
	ResponsibleID   string `json:"responsible_id,omitempty"`
	ResponsibleName string `json:"responsible_name,omitempty"`
	EmisorID        string `json:"emisor_id,omitempty"`
	EmisorName      string `json:"emisor_name,omitempty"`
}

// ClientAssignmentMatrixRow represents a row in the client-assignment matrix for the frontend
// Each row is a client-assignment pair with a boolean indicating if the client has that assignment type
// Useful for building a dynamic table in the frontend
type ClientAssignmentMatrixRow struct {
	ClientID           string `json:"client_id"`
	ClientName         string `json:"client_name"`
	AssignmentTypeID   int    `json:"assignment_type_id"`
	AssignmentTypeName string `json:"assignment_type_name"`
	Selected           bool   `json:"selected"`
}

type AssignmentSelection struct {
	AssignmentTypeID int  `json:"assignment_type_id"`
	Selected         bool `json:"selected"`
}

// ACCOUNTANCY STATUS

type ClientAccountancyStatus struct {
	ID          int    `json:"id"`
	ClientID    string `json:"client_id"`
	Month       string `json:"month"`              // YYYY-MM-DD
	DueDate     string `json:"due_date,omitempty"` // YYYY-MM-DD
	Observacion string `json:"observaciones,omitempty"`
}

type ClientAccountancyAssignment struct {
	ID                   int    `json:"id"`
	StatusID             int    `json:"status_id"`
	AssignmentTypeID     int    `json:"assignment_type_id"`
	AssignmentTypeName   string `json:"assignment_type_name"`
	AssignmentStatusID   int    `json:"assignment_status_id"`
	AssignmentStatusName string `json:"assignment_status_name"`
}

type ClientAccountancyHistoryEntry struct {
	Status      ClientAccountancyStatus       `json:"status"`
	Assignments []ClientAccountancyAssignment `json:"assignments"`
}

// Estructura para devolver el historial y los tipos de asignaciones activas

type ClientAccountancyHistoryWithAssignments struct {
	History           []ClientAccountancyHistoryEntry `json:"history"`
	ActiveAssignments []AccountancyType               `json:"active_assignments"`
}

type UpdateClientResponsible struct {
	ResponsibleID string `json:"responsible_id"`
	ClientID      string `json:"client_id"`
}
