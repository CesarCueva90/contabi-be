package usecase

import "contabi-be/models"

// NominasInteractor implements the NominasUseCase interface
type NominasInteractor struct {
	nominasService NominasService
}

// NewNominasUseCase creates a new instance of NominasService
func NewNominasUseCase(nominasService NominasService) NominasService {
	return &NominasInteractor{
		nominasService: nominasService,
	}
}

// CreateClientPaymentRecord creates a new client payment record
func (ni *NominasInteractor) CreateClientPaymentRecord(clientPaymentRecord models.ClientHRPayment) error {
	return ni.nominasService.CreateClientPaymentRecord(clientPaymentRecord)
}

// CreateClientPaymentRecord creates a new client payment record
func (ni *NominasInteractor) GetClientsWithPendingPaymentsByHREntityID(hrEntityID string) ([]models.ClientWithPendingHRPayment, error) {
	return ni.nominasService.GetClientsWithPendingPaymentsByHREntityID(hrEntityID)
}

// CreateClientPaymentRecord creates a new client payment record
func (ni *NominasInteractor) GetClientPendingPaymentsByHREntityIDDetails(clientID, hrEntityID string) ([]models.ClientWithPendingHRPaymentDetails, error) {
	return ni.nominasService.GetClientPendingPaymentsByHREntityIDDetails(clientID, hrEntityID)
}

// UpdateClientPaymentRecord updates a client payment record
func (ni *NominasInteractor) UpdateClientPaymentRecord(clientPaymentRecord models.UpdateClientHRPayment) error {
	return ni.nominasService.UpdateClientPaymentRecord(clientPaymentRecord)
}

// GetClientHRPaymentsHistory gets all the payments of a specific client and HR entity
func (ni *NominasInteractor) GetClientHRPaymentsHistory(clientID, hrEntityID string) ([]models.ClientHRPayment, error) {
	return ni.nominasService.GetClientHRPaymentsHistory(clientID, hrEntityID)
}
