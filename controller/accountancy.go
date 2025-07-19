package controller

import (
	"contabi-be/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AccountancyController handles basic CRUD operations for clients
type AccountancyController struct {
	accountancyUseCase AccountancyUseCase
	logger             *logrus.Logger
}

func NewAccountancyController(accountancyUseCase AccountancyUseCase, logger *logrus.Logger) *AccountancyController {
	return &AccountancyController{
		accountancyUseCase: accountancyUseCase,
		logger:             logger,
	}
}

// GetClientsBySupervisor retrieves all the clients of a specific supervisor
func (ac *AccountancyController) GetClientsBySupervisor(g *gin.Context) {
	supervisorID := g.Param("supervisor_id")
	clients, err := ac.accountancyUseCase.GetClientsBySupervisor(supervisorID)
	if err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetClientsBySupervisor(): Error fetching clients info")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching clients info"})
		return
	}

	g.JSON(http.StatusOK, clients)
}

// GetClientAssignmentsMatrix retrieves, for all active clients, the list of assignment types and whether each client has each assignment
func (ac *AccountancyController) GetClientAssignmentsMatrix(g *gin.Context) {
	assignments, err := ac.accountancyUseCase.GetClientAssignmentsMatrix()
	if err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetClientAssignmentsMatrix(): Error fetching clients accountancy assignments")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching clients accountancy assignments"})
		return
	}

	g.JSON(http.StatusOK, assignments)
}

// UpdateClientAssignments updates assignments of a client according to the state
func (ac *AccountancyController) UpdateClientAssignments(g *gin.Context) {
	clientID := g.Param("client_id")

	var assignments []models.AssignmentSelection
	if err := g.ShouldBindJSON(&assignments); err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding client data")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Error binding client data"})
		return
	}

	err := ac.accountancyUseCase.UpdateClientAssignments(clientID, assignments)
	if err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("UpdateClientAssignments(): error updating client accountancy assignments")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating client accountancy assignments"})
		return
	}

	g.JSON(http.StatusOK, "Client accountancy assignments updated successfully")
}

// GetClientsBySResonsible retrieves all the clients of a specific responsible
func (ac *AccountancyController) GetClientsByResonsible(g *gin.Context) {
	supervisorID := g.Param("responsible_id")
	clients, err := ac.accountancyUseCase.GetClientsByResonsible(supervisorID)
	if err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetClientsByResonsible(): Error fetching clients info")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching clients info"})
		return
	}

	g.JSON(http.StatusOK, clients)
}

// CreateClientAccountancyStatusWithAssignments creates a new monthly record for a client
func (ac *AccountancyController) CreateClientAccountancyStatusWithAssignments(g *gin.Context) {
	var req models.ClientAccountancyHistoryEntry

	if err := g.ShouldBindJSON(&req); err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding accountancy status data")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Error binding accountancy status data"})
		return
	}

	err := ac.accountancyUseCase.CreateClientAccountancyStatusWithAssignments(req.Status, req.Assignments)
	if err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("CreateClientAccountancyStatusWithAssignments(): error creating accountancy status and assignments")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating accountancy status and assignments"})
		return
	}

	g.JSON(http.StatusCreated, "Client accountancy status and assignments created successfully")
}

// GetClientAccountancyHistory gets the hisotory for a client accountancy behavior
func (ac *AccountancyController) GetClientAccountancyHistory(g *gin.Context) {
	clientID := g.Param("client_id")
	result, err := ac.accountancyUseCase.GetClientAccountancyHistory(clientID)
	if err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetClientAccountancyHistory(): Error fetching clients info")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching clients info"})
		return
	}

	g.JSON(http.StatusOK, result)
}

// UpdateClientAccountancyStatusWithAssignments updates an existing monthly record for a client
func (ac *AccountancyController) UpdateClientAccountancyStatusWithAssignments(g *gin.Context) {
	clientID := g.Param("client_id")
	statusID := g.Param("status_id")

	var req models.ClientAccountancyHistoryEntry
	if err := g.ShouldBindJSON(&req); err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding accountancy status update data")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Error binding accountancy status update data"})
		return
	}

	// Validar que el status pertenece al cliente especificado
	if req.Status.ClientID != clientID {
		ac.logger.WithFields(logrus.Fields{
			"client_id":        clientID,
			"status_client_id": req.Status.ClientID,
		}).Error("Status does not belong to the specified client")
		g.JSON(http.StatusForbidden, gin.H{"error": "Status does not belong to this client"})
		return
	}

	// Convertir statusID a int
	var statusIDInt int
	if _, err := fmt.Sscanf(statusID, "%d", &statusIDInt); err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error parsing status_id")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status_id"})
		return
	}

	err := ac.accountancyUseCase.UpdateClientAccountancyStatusWithAssignments(statusIDInt, clientID, req.Status, req.Assignments)
	if err != nil {
		ac.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("UpdateClientAccountancyStatusWithAssignments(): error updating accountancy status and assignments")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating accountancy status and assignments"})
		return
	}

	g.JSON(http.StatusOK, "Client accountancy status and assignments updated successfully")
}
