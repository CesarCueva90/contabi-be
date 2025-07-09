package controller

import (
	"contabi-be/models"

	"github.com/sirupsen/logrus"
)

// LoginUseCase
type LoginUseCase interface {
	Login(login, password string) (models.User, error)
}

// UsersUseCase
type UsersUseCase interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	UpdateUserRole(user models.User) error
	PutUserPassword(user models.User) error
	DeleteUser(id string) error
	GetRoles() ([]models.Role, error)
}

type ClientsUsecase interface {
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

type MenusUseCase interface {
	GetEmisors() ([]models.Emisor, error)
	GetSupervisors() ([]models.Supervisor, error)
	GetResponsiblesBySupervisor(supervisorID string) ([]models.Responsible, error)
	GetRegimenes() ([]models.Regimen, error)
}

// Controller
type Controller struct {
	lu     LoginUseCase
	uu     UsersUseCase
	cu     ClientsUsecase
	mu     MenusUseCase
	logger *logrus.Logger
}

// NewController creates a new Controllert instance
func NewController(lu LoginUseCase, uu UsersUseCase, cu ClientsUsecase, mu MenusUseCase, logger *logrus.Logger) *Controller {
	return &Controller{
		lu:     lu,
		uu:     uu,
		cu:     cu,
		mu:     mu,
		logger: logger,
	}
}
