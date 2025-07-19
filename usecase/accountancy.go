package usecase

import "contabi-be/models"

// AccountancyInteractor implements the MenusUseCase interface
type AccountancyInteractor struct {
	accountancyService AccountancyService
}

// NewAccountancyUseCase creates a new instance of AccountancyService
func NewAccountancyUseCase(accountancyService AccountancyService) AccountancyService {
	return &AccountancyInteractor{
		accountancyService: accountancyService,
	}
}

// GetClientsBySupervisor retrieves all the clients of a specific supervisor
func (ai *AccountancyInteractor) GetClientsBySupervisor(supervisorID string) ([]models.AccountancyClientInfo, error) {
	return ai.accountancyService.GetClientsBySupervisor(supervisorID)
}

// GetClientAssignmentsMatrix retrieves, for all active clients, the list of assignment types and whether each client has each assignment
func (ai *AccountancyInteractor) GetClientAssignmentsMatrix() ([]models.ClientAssignmentMatrixRow, error) {
	return ai.accountancyService.GetClientAssignmentsMatrix()
}

// UpdateClientAssignments updates assignments of a client according to the state
func (ai *AccountancyInteractor) UpdateClientAssignments(clientID string, assignments []models.AssignmentSelection) error {
	return ai.accountancyService.UpdateClientAssignments(clientID, assignments)
}

// GetClientsBySResonsible retrieves all the clients of a specific responsible
func (ai *AccountancyInteractor) GetClientsByResonsible(responsibleID string) ([]models.AccountancyClientInfo, error) {
	return ai.accountancyService.GetClientsByResonsible(responsibleID)
}

// CreateClientAccountancyStatusWithAssignments creates a new monthly record for a client
func (ai *AccountancyInteractor) CreateClientAccountancyStatusWithAssignments(status models.ClientAccountancyStatus, assignments []models.ClientAccountancyAssignment) error {
	return ai.accountancyService.CreateClientAccountancyStatusWithAssignments(status, assignments)
}

// GetClientAccountancyHistory gets the hisotory for a client accountancy behavior
func (ai *AccountancyInteractor) GetClientAccountancyHistory(clientID string) (models.ClientAccountancyHistoryWithAssignments, error) {
	return ai.accountancyService.GetClientAccountancyHistory(clientID)
}

// UpdateClientAccountancyStatusWithAssignments updates an existing monthly record for a client
func (ai *AccountancyInteractor) UpdateClientAccountancyStatusWithAssignments(statusID int, clientID string, status models.ClientAccountancyStatus, assignments []models.ClientAccountancyAssignment) error {
	return ai.accountancyService.UpdateClientAccountancyStatusWithAssignments(statusID, clientID, status, assignments)
}
