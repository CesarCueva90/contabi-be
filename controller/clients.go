package controller

import (
	"contabi-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ClientsController handles basic CRUD operations for clients
type ClientsController struct {
	clientsUseCase ClientsUsecase
	logger         *logrus.Logger
}

func NewClientsController(clientsUseCase ClientsUsecase, logger *logrus.Logger) *ClientsController {
	return &ClientsController{
		clientsUseCase: clientsUseCase,
		logger:         logger,
	}
}

// GetClientsInfo returns all clients with complete information
func (cc *ClientsController) GetClientsInfo(c *gin.Context) {
	clients, err := cc.clientsUseCase.GetAllClientsInfo()
	if err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetAllClientsInfo(): Error fetching clients info")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching clients info"})
		return
	}

	c.JSON(http.StatusOK, clients)
}

// GetActiveClientsInfo returns only active clients with complete information
func (cc *ClientsController) GetActiveClientsInfo(c *gin.Context) {
	clients, err := cc.clientsUseCase.GetActiveClientsInfo()
	if err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetActiveClientsInfo(): Error fetching active clients info")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching active clients info"})
		return
	}

	c.JSON(http.StatusOK, clients)
}

// GetClientInfo returns complete information for a specific client
func (cc *ClientsController) GetClientInfo(c *gin.Context) {
	clientID := c.Param("id")

	client, err := cc.clientsUseCase.GetClientInfo(clientID)
	if err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetClientInfo(): Error fetching client info")
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching client info"})
		return
	}

	c.JSON(http.StatusOK, client)
}

// CreateClient creates a new client with assignments
func (cc *ClientsController) CreateClient(c *gin.Context) {
	var request struct {
		Client      models.Client            `json:"client"`
		Assignments models.ClientAssignments `json:"assignments"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding client data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding client data"})
		return
	}

	err := cc.clientsUseCase.CreateClient(request.Client, request.Assignments)
	if err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("CreateClient(): Error creating client")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating client"})
		return
	}

	c.JSON(http.StatusCreated, "Client created successfully")
}

// UpdateClient updates client basic information
func (cc *ClientsController) UpdateClient(c *gin.Context) {
	clientID := c.Param("id")

	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding client data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding client data"})
		return
	}

	err := cc.clientsUseCase.UpdateClient(clientID, client)
	if err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("UpdateClient(): error updating client")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating client"})
		return
	}

	c.JSON(http.StatusOK, "Client updated successfully")
}

// DeactivateClient sets a client as inactive (soft delete)
func (cc *ClientsController) DeactivateClient(c *gin.Context) {
	clientID := c.Param("id")

	err := cc.clientsUseCase.DeactivateClient(clientID)
	if err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("DeactivateClient(): error while deleting client")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting client"})
		return
	}

	c.JSON(http.StatusOK, "Client deactivated successfully")
}

// ActivateClient sets a client as active
func (cc *ClientsController) ActivateClient(c *gin.Context) {
	clientID := c.Param("id")

	err := cc.clientsUseCase.ActivateClient(clientID)
	if err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("ActivateClient(): error while activating client")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while activating client"})
		return
	}

	c.JSON(http.StatusOK, "Client activated successfully")
}

// UpdateClientAssignments updates client assignments (supervisor, responsible, emisor)
func (cc *ClientsController) UpdateClientAssignments(c *gin.Context) {
	clientID := c.Param("id")

	var assignments models.ClientAssignments
	if err := c.ShouldBindJSON(&assignments); err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding client data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding client data"})
		return
	}

	err := cc.clientsUseCase.UpdateClientAssignments(clientID, assignments)
	if err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("UpdateClientAssignments(): error while updating client assignments")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating client assignments"})
		return
	}

	c.JSON(http.StatusOK, "Client assignments updated successfully")
}

// GetClientsWithPendingPayments returns clients that have pending payments
func (cc *ClientsController) GetClientsWithPendingPayments(c *gin.Context) {
	clients, err := cc.clientsUseCase.GetClientsWithPendingPayments()
	if err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetClientsWithPendingPayments(): error while getting clients payments")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while getting clients payments"})
		return
	}

	c.JSON(http.StatusOK, clients)
}

// UpdateClientPayment updates client payment information
func (cc *ClientsController) UpdateClientPayment(c *gin.Context) {
	clientID := c.Param("id")

	var payment models.ClientPayment
	if err := c.ShouldBindJSON(&payment); err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding client data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding client data"})
		return
	}

	err := cc.clientsUseCase.UpdateClientPayment(clientID, payment)
	if err != nil {
		cc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("UpdateClientPayment(): error while updating client payment")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Client payment updated successfully")
}
