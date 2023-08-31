package services

import (
	"context"

	"github.com/NithishNithi/netxd_grpc_mongo_dal/netxd_grpc_mongo_dal/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection1 *mongo.Collection
	CustomerCollection2 *mongo.Collection
	ctx                 context.Context
}

// InitCustomerService initializes a CustomerService instance.
func InitCustomerService(collection1 *mongo.Collection, collection2 *mongo.Collection, ctx context.Context) interfaces.ICustomers {
	return &CustomerService{collection1, collection2, ctx}
}
