package usecase

import "contabi-be/models"

// MenusInteractor implements the MenusUseCase interface
type MenusInteractor struct {
	menusService MenusService
}

// NewMenusUseCase creates a new instance of MenusService
func NewMenusUseCase(menusService MenusService) MenusService {
	return &MenusInteractor{
		menusService: menusService,
	}
}

// GetEmisors retrieves all emisors
func (mi *MenusInteractor) GetEmisors() ([]models.Emisor, error) {
	return mi.menusService.GetEmisors()
}

// GetSupervisors retrieves all supervisors
func (mi *MenusInteractor) GetSupervisors() ([]models.Supervisor, error) {
	return mi.menusService.GetSupervisors()
}

// GetResponsiblesBySupervisor retrieves all responsibles
func (mi *MenusInteractor) GetResponsiblesBySupervisor(supervisorID string) ([]models.Responsible, error) {
	return mi.menusService.GetResponsiblesBySupervisor(supervisorID)
}

// GetRegimenes retrieves all regimenes
func (mi *MenusInteractor) GetRegimenes() ([]models.Regimen, error) {
	return mi.menusService.GetRegimenes()
}

// GetAccountancyTypes retrieves all accountancy types
func (mi *MenusInteractor) GetAccountancyTypes() ([]models.AccountancyType, error) {
	return mi.menusService.GetAccountancyTypes()
}

// GetAccountancyStatuses retrieves all accountancy assignment statuses
func (mi *MenusInteractor) GetAccountancyStatuses() ([]models.AccountancyAssignmentStatus, error) {
	return mi.menusService.GetAccountancyStatuses()
}
