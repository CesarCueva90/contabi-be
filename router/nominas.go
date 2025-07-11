package router

import "github.com/gin-gonic/gin"

func nominasRouter(r *gin.Engine, nominasController NominasController) {
	r.POST("/client/hrpayment", nominasController.CreateClientPaymentRecord)
	r.GET("/clients/hrpayment/:hr_entity_id", nominasController.GetClientsWithPendingPaymentsByHREntityID)
	r.GET("/client/:client_id/hrpayments/:hr_entity_id", nominasController.GetClientPendingPaymentsByHREntityIDDetails)
	r.PUT("/client/hrpayment", nominasController.UpdateClientPaymentRecord)
	r.GET("/client/:client_id/hrpayments/:hr_entity_id/history", nominasController.GetClientHRPaymentsHistory)
}
