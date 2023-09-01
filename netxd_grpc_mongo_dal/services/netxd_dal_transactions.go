package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/NithishNithi/netxd_grpc_mongo_dal/netxd_grpc_mongo_dal/models"
	"go.mongodb.org/mongo-driver/bson"
)

var Customer1 struct {
	CustomerId int     `json:"customerid" bson:"customerid"`
	Balance    float64 `json:"balance" bson:"balance"`
	
}
var Customer2 struct {
	CustomerId int     `json:"customerid" bson:"customerid"`
	Balance    float64 `json:"balance" bson:"balance"`
	
}

func (p *CustomerService) MakeTransactions(transaction *models.MakeTransactions) (*models.UpdateResponse, error) {
	filter1 := bson.M{"customerid": transaction.From_CustomerId}
	filter2 := bson.M{"customerid": transaction.To_CustomerId}

	err1 := p.CustomerCollection1.FindOne(context.Background(), filter1).Decode(&Customer1)
	if err1 != nil {
		log.Fatal(err1)
		return nil, err1
	}
	err2 := p.CustomerCollection1.FindOne(context.Background(), filter2).Decode(&Customer2)
	if err1 != nil {
		log.Fatal(err2)
		return nil, err2
	}
	if transaction.Amount <= Customer1.Balance && transaction.Amount > 0 {
		Customer2.Balance += transaction.Amount
		Customer1.Balance -= transaction.Amount
	} else {
		panic("Insufficient Funds ")
	}

	cus1_filter := bson.M{"customer_id": transaction.From_CustomerId}
	cus2_filter := bson.M{"customer_id": transaction.To_CustomerId}

	cus1_update := bson.M{"$set": bson.M{"balance_amount": Customer1.Balance}}
	cus2_update := bson.M{"$set": bson.M{"balance_amount": Customer2.Balance}}

	_, cus1_err := p.CustomerCollection1.UpdateOne(context.Background(), cus1_filter, cus1_update)
	if cus1_err != nil {
		log.Fatal(cus1_err)
		return nil, cus1_err
	}
	_, cus2_err := p.CustomerCollection1.UpdateOne(context.Background(), cus2_filter, cus2_update)
	if cus1_err != nil {
		log.Fatal(cus2_err)
		return nil, cus2_err
	}

	t := time.Now()
	TransTime := t.Format(time.Kitchen)

	TransactionRecord := models.MakeTransactions{
		From_CustomerId: transaction.From_CustomerId,
		To_CustomerId:   transaction.To_CustomerId,
		Amount:          transaction.Amount,
		Transactiontime: TransTime,
	}
	res, transErr := p.CustomerCollection2.InsertOne(context.Background(), TransactionRecord)
	if transErr != nil {
		fmt.Println("Error while inserting into transactions Table", transErr)
		return nil, transErr
	} else {
		fmt.Println("Insertion Successfully Done")
	}

	newUser := &models.UpdateResponse{}
	query := bson.M{"_id": res.InsertedID}

	err := p.CustomerCollection1.FindOne(p.ctx, query).Decode(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
