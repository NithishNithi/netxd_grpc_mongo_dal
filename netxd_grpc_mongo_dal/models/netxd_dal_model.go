package models

type Customers struct {
	CustomerId int     `json:"customerid" bson:"customerid"`
	FirstName  string  `json:"firstname" bson:"firstname"`
	LastName   string  `json:"lastname" bson:"lastname"`
	BankId     string  `json:"bankid" bson:"bankid"`
	Balance    float64 `json:"balance" bson:"balance"`
	CreatedAt  string  `json:"createdat" bson:"createdat"`
	UpdatedAt  string  `json:"updatedat" bson:"updatedat"`
	IsActive   string  `json:"isactive" bson:"isactive"`
}

type CustomerResponse struct {
	CustomerId int    `json:"customerid" bson:"customerid"`
	CreatedAt  string `json:"isactive" bson:"isactive"`
}

type MakeTransactions struct {
	From_CustomerId int	 `json:"fromcustomerid" bson:"fromcustomerid"`
	To_CustomerId   int	`json:"tocustomerid" bson:"tocustomerid"`
	Amount          float64	`json:"amount" bson:"amount"`
	Transactiontime string	`json:"transactiontime" bson:"transactiontime"`
}
type UpdateResponse struct {
	From_CustomerId_Balance float64	`json:"fromcustomeridbalance" bson:"fromcustomeridbalance"`
	To_CustomerId_Balance float64	`json:"tocustomeridbalance" bson:"tocustomeridbalance"`
	Transactiontime string	`json:"transactiontime" bson:"transactiontime"`
}
