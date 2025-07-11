package controller

import (
	"contabi-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NominasController handles basic CRUD operations for clients
type NominasController struct {
	nominasUsecase NominasUseCase
	logger         *logrus.Logger
}

func NewNominasController(nominasUsecase NominasUseCase, logger *logrus.Logger) *NominasController {
	return &NominasController{
		nominasUsecase: nominasUsecase,
		logger:         logger,
	}
}

// CreateClientPaymentRecord creates a new client payment record
func (nc *NominasController) CreateClientPaymentRecord(g *gin.Context) {
	var clientPaymentRecord models.ClientHRPayment
	if err := g.ShouldBindJSON(&clientPaymentRecord); err != nil {
		nc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding data")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := nc.nominasUsecase.CreateClientPaymentRecord(clientPaymentRecord)
	if err != nil {
		nc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("CreateClientPaymentRecord(): error while creating client payment record")
		g.JSON(http.StatusBadRequest, gin.H{"error": "error while creating client payment record"})
		return
	}

	g.JSON(http.StatusCreated, "client payment record created successfully")
}

// GetClientsWithPendingPaymentsByHREntityID gets the list of all clients with pending payments of specific HR entity
func (nc *NominasController) GetClientsWithPendingPaymentsByHREntityID(g *gin.Context) {
	hrEntityID := g.Param("hr_entity_id")

	clients, err := nc.nominasUsecase.GetClientsWithPendingPaymentsByHREntityID(hrEntityID)
	if err != nil {
		nc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error fetching clients with pending payments")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching clients with pending payments"})
		return
	}

	g.JSON(http.StatusOK, clients)
}

// GetClientPendingPaymentsByHREntityIDDetails gets all the pending payments of a specific client and HR entity
func (nc *NominasController) GetClientPendingPaymentsByHREntityIDDetails(g *gin.Context) {
	clientID := g.Param("client_id")
	hrEntityID := g.Param("hr_entity_id")

	payments, err := nc.nominasUsecase.GetClientPendingPaymentsByHREntityIDDetails(clientID, hrEntityID)
	if err != nil {
		nc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error fetching pending payments of the client")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching pending payments of the client"})
		return
	}

	g.JSON(http.StatusOK, payments)
}

// UpdateClientPaymentRecord updates a client payment record
func (nc *NominasController) UpdateClientPaymentRecord(g *gin.Context) {
	var clientPaymentRecord models.UpdateClientHRPayment
	if err := g.ShouldBindJSON(&clientPaymentRecord); err != nil {
		nc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding data")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := nc.nominasUsecase.UpdateClientPaymentRecord(clientPaymentRecord)
	if err != nil {
		nc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("UpdateClientPaymentRecord(): error while updating client payment record")
		g.JSON(http.StatusBadRequest, gin.H{"error": "error while updating client payment record"})
		return
	}

	g.JSON(http.StatusOK, "client payment record updating successfully")
}

// GetClientHRPaymentsHistory gets all the payments of a specific client and HR entity
func (nc *NominasController) GetClientHRPaymentsHistory(g *gin.Context) {
	clientID := g.Param("client_id")
	hrEntityID := g.Param("hr_entity_id")

	payments, err := nc.nominasUsecase.GetClientHRPaymentsHistory(clientID, hrEntityID)
	if err != nil {
		nc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error fetching history payments of the client")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching history payments of the client"})
		return
	}

	g.JSON(http.StatusOK, payments)
}
