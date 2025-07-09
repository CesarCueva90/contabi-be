package usecase

import "contabi-be/models"

// LoginUseCase defines the interface for login-related operations
type LoginUseCase interface {
	Login(login, password string) (models.User, error)
}

// LoginService defines the interface for login-related operations
type LoginService interface {
	Login(login, password string) (models.User, error)
}

type UsersService interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	UpdateUserRole(user models.User) error
	PutUserPassword(user models.User) error
	DeleteUser(id string) error
	GetRoles() ([]models.Role, error)
}

// ClientsService defines the interface for client CRUD operations
type ClientsService interface {
	GetAllClientsInfo() ([]models.ClientInfo, error)
	GetActiveClientsInfo() ([]models.ClientInfo, error)
	GetClientInfo(clientID string) (models.ClientInfo, error)
	CreateClient(client models.Client, assignments models.ClientAssignments) error
	UpdateClient(clientID string, client models.Client) error
	DeactivateClient(clientID string) error
	ActivateClient(clientID string) error
	UpdateClientAssignments(clientID string, assignments models.ClientAssignments) error
	GetClientsWithPendingPayments() ([]models.ClientWithPendingPayment, error)
	UpdateClientPayment(clientID string, payment models.ClientPayment) error
	GetClientPayments(clientID string) ([]models.ClientPaymentHistory, error)
}

type MenusService interface {
	GetEmisors() ([]models.Emisor, error)
	GetSupervisors() ([]models.Supervisor, error)
	GetResponsiblesBySupervisor(supervisorID string) ([]models.Responsible, error)
	GetRegimenes() ([]models.Regimen, error)
}
