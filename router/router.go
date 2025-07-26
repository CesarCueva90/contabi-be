package router

import (
	"contabi-be/middleware"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(g *gin.Context)
}

type UsersController interface {
	GetUsers(g *gin.Context)
	GetUserByID(g *gin.Context)
	CreateUser(g *gin.Context)
	UpdateUser(g *gin.Context)
	UpdateUserRole(g *gin.Context)
	PutUserPassword(g *gin.Context)
	DeleteUser(g *gin.Context)
	GetRoles(g *gin.Context)
}

// ClientsController handles all client operations
type ClientsController interface {
	GetClientsInfo(c *gin.Context)
	GetActiveClientsInfo(c *gin.Context)
	GetClientInfo(c *gin.Context)
	CreateClient(c *gin.Context)
	UpdateClient(c *gin.Context)
	DeactivateClient(c *gin.Context)
	ActivateClient(c *gin.Context)
	UpdateClientAssignments(c *gin.Context)
	GetClientsWithPendingPayments(c *gin.Context)
	UpdateClientPayment(c *gin.Context)
	GetClientPayments(c *gin.Context)
}

type MenusController interface {
	GetEmisors(g *gin.Context)
	GetSupervisors(g *gin.Context)
	GetResponsiblesBySupervisor(g *gin.Context)
	GetRegimenes(g *gin.Context)
	GetAccountancyTypes(g *gin.Context)
	GetAccountancyStatuses(g *gin.Context)
}

type NominasController interface {
	CreateClientPaymentRecord(g *gin.Context)
	GetClientsWithPendingPaymentsByHREntityID(g *gin.Context)
	GetClientPendingPaymentsByHREntityIDDetails(g *gin.Context)
	UpdateClientPaymentRecord(g *gin.Context)
	GetClientHRPaymentsHistory(g *gin.Context)
}

type AccountancyController interface {
	GetClientsBySupervisor(g *gin.Context)
	GetClientAssignmentsMatrix(g *gin.Context)
	UpdateClientAssignments(g *gin.Context)
	GetClientsByResonsible(g *gin.Context)
	CreateClientAccountancyStatusWithAssignments(g *gin.Context)
	UpdateClientAccountancyStatusWithAssignments(g *gin.Context)
	GetClientAccountancyHistory(g *gin.Context)
	GetAllClients(g *gin.Context)
}

// NewRouter set the API routes and applies the middleware
func NewRouter(
	loginController LoginController,
	usersController UsersController,
	clientsController ClientsController,
	menusController MenusController,
	nominasController NominasController,
	accountancyController AccountancyController,
	mw *middleware.Middleware,
) *gin.Engine {
	// Creates a new instance of Gin router
	r := gin.Default()

	// Adds the CORS middleware to all routes
	r.Use(mw.CORS())

	// Routes for Login
	loginRoutes(r, loginController)

	// Adds the authentication middleware to the required routes
	r.Use(mw.AuthMiddleware())

	usersRoutes(r, usersController)

	clientsRoutes(r, clientsController)

	menusRoutes(r, menusController)

	nominasRouter(r, nominasController)

	accountancyRoutes(r, accountancyController)

	return r
}
