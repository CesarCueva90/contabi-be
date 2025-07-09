package router

import (
	"github.com/gin-gonic/gin"
)

func clientsRoutes(r *gin.Engine, clientsController ClientsController) {
	// Gets all clients with full info
	r.GET("/clients", clientsController.GetClientsInfo)

	// Gets only active clients with full info
	r.GET("/clients/active", clientsController.GetActiveClientsInfo)

	// Gets full info of a specific client
	r.GET("/clients/:id", clientsController.GetClientInfo)

	// Creates a new client with assignments
	r.POST("/clients/", clientsController.CreateClient)

	// Updates the basic info of a client
	r.PUT("/clients/:id", clientsController.UpdateClient)

	// Deactivates a client (soft delete)
	r.DELETE("/clients/:id", clientsController.DeactivateClient)

	// Activates a client
	r.PUT("/clients/:id/activate", clientsController.ActivateClient)

	// Updates the assignments of a specific client (supervisor, responsible, emisor)
	r.PUT("/clients/:id/assignments", clientsController.UpdateClientAssignments)

	// Get clients with pending payments
	r.GET("/clients/pending-payments", clientsController.GetClientsWithPendingPayments)

	// Updates the payment info of a specific client
	r.PUT("/clients/:id/payment", clientsController.UpdateClientPayment)
}
