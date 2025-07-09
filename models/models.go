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
