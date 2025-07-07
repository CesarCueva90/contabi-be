package router

import (
	"github.com/gin-gonic/gin"
)

func clientsRoutes(r *gin.Engine, clientsController ClientsController) {
	clients := r.Group("/clients")
	{
		// Gets all clients with full info
		clients.GET("/", clientsController.GetClientsInfo)

		// Gets only active clients with full info
		clients.GET("/active", clientsController.GetActiveClientsInfo)

		// Gets full info of a specific client
		clients.GET("/:id", clientsController.GetClientInfo)

		// Crear un nuevo cliente con asignaciones
		clients.POST("/", clientsController.CreateClient)

		// Actualizar información básica del cliente
		clients.PUT("/:id", clientsController.UpdateClient)

		// Desactivar cliente (soft delete)
		clients.DELETE("/:id", clientsController.DeactivateClient)

		// Activar cliente
		clients.PUT("/:id/activate", clientsController.ActivateClient)

		// Actualizar asignaciones del cliente (supervisor, responsable, emisor)
		clients.PUT("/:id/assignments", clientsController.UpdateClientAssignments)

		// Get clients with pending payments
		clients.GET("/pending-payments", clientsController.GetClientsWithPendingPayments)

		// Actualizar información de pago del cliente
		clients.PUT("/:id/payment", clientsController.UpdateClientPayment)
	}
}
