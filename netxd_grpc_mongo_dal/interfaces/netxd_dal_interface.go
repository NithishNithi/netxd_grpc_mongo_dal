package interfaces

import (
	"github.com/NithishNithi/netxd_grpc_mongo_dal/netxd_grpc_mongo_dal/models"
)

type ICustomers interface {
	CreateCustomer(customer *models.Customers) (*models.CustomerResponse, error)
	MakeTransactions(transaction *models.MakeTransactions)(*models.UpdateResponse,error)
}
