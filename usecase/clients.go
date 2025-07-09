package usecase

import (
	"contabi-be/models"
)

// ClientsInteractor implements the ClientsService interface
type ClientsInteractor struct {
	clientsService ClientsService
}

// NewClientsUseCase creates a new instance of ClientsService
func NewClientsUseCase(clientsService ClientsService) ClientsService {
	return &ClientsInteractor{
		clientsService: clientsService,
	}
}

// GetAllClientsInfo retrieves all clients with complete information
func (ci *ClientsInteractor) GetAllClientsInfo() ([]models.ClientInfo, error) {
	return ci.clientsService.GetAllClientsInfo()
}

// GetActiveClientsInfo retrieves only active clients with complete information
func (ci *ClientsInteractor) GetActiveClientsInfo() ([]models.ClientInfo, error) {
	return ci.clientsService.GetActiveClientsInfo()
}

// GetClientInfo retrieves complete information for a specific client
func (ci *ClientsInteractor) GetClientInfo(clientID string) (models.ClientInfo, error) {
	return ci.clientsService.GetClientInfo(clientID)
}

// CreateClient creates a new client with assignments
func (ci *ClientsInteractor) CreateClient(client models.Client, assignments models.ClientAssignments) error {
	return ci.clientsService.CreateClient(client, assignments)
}

// UpdateClient updates client basic information
func (ci *ClientsInteractor) UpdateClient(clientID string, client models.Client) error {
	return ci.clientsService.UpdateClient(clientID, client)
}

// DeactivateClient sets a client as inactive (soft delete)
func (ci *ClientsInteractor) DeactivateClient(clientID string) error {
	return ci.clientsService.DeactivateClient(clientID)
}

// ActivateClient sets a client as active
func (ci *ClientsInteractor) ActivateClient(clientID string) error {
	return ci.clientsService.ActivateClient(clientID)
}

// UpdateClientAssignments updates client assignments (supervisor, responsible, emisor)
func (ci *ClientsInteractor) UpdateClientAssignments(clientID string, assignments models.ClientAssignments) error {
	return ci.clientsService.UpdateClientAssignments(clientID, assignments)
}

// GetClientsWithPendingPayments returns clients that have pending payments
func (ci *ClientsInteractor) GetClientsWithPendingPayments() ([]models.ClientWithPendingPayment, error) {
	return ci.clientsService.GetClientsWithPendingPayments()
}

// UpdateClientPayment updates client payment information
func (ci *ClientsInteractor) UpdateClientPayment(clientID string, payment models.ClientPayment) error {
	return ci.clientsService.UpdateClientPayment(clientID, payment)
}

// GetClientPayments gets the payments history of a specific client
func (ci *ClientsInteractor) GetClientPayments(clientID string) ([]models.ClientPaymentHistory, error) {
	return ci.clientsService.GetClientPayments(clientID)
}
