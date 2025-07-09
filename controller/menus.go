package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// MenusController handles the user-related HTTP requests
type MenusController struct {
	menusUseCase MenusUseCase
	logger       *logrus.Logger
}

// NewMenusController creates a new instance of MenusController
func NewMenusController(menusUseCase MenusUseCase, logger *logrus.Logger) *MenusController {
	return &MenusController{
		menusUseCase: menusUseCase,
		logger:       logger,
	}
}

// GetEmisors handles the request to fetch all emisors
func (mc *MenusController) GetEmisors(g *gin.Context) {
	emisors, err := mc.menusUseCase.GetEmisors()
	if err != nil {
		mc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetEmisors(): Error fetching emisors")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching emisors"})
		return
	}

	g.JSON(http.StatusOK, emisors)
}

// GetSupervisors handles the request to fetch all supervisors
func (mc *MenusController) GetSupervisors(g *gin.Context) {
	supervisors, err := mc.menusUseCase.GetSupervisors()
	if err != nil {
		mc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetSupervisors(): Error fetching supervisors")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching supervisors"})
		return
	}

	g.JSON(http.StatusOK, supervisors)
}

// GetResponsiblesBySupervisor retrieves all responsibles
func (mc *MenusController) GetResponsiblesBySupervisor(g *gin.Context) {
	supervisorID := g.Param("supervisor_id")
	responsibles, err := mc.menusUseCase.GetResponsiblesBySupervisor(supervisorID)
	if err != nil {
		mc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetResponsiblesBySupervisor(): Error fetching responsibles")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching responsibles"})
		return
	}

	g.JSON(http.StatusOK, responsibles)
}

// GetRegimenes retrieves all regimenes
func (mc *MenusController) GetRegimenes(g *gin.Context) {
	regimenes, err := mc.menusUseCase.GetRegimenes()
	if err != nil {
		mc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetRegimenes(): Error fetching regimenes")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching regimenes"})
		return
	}

	g.JSON(http.StatusOK, regimenes)
}
