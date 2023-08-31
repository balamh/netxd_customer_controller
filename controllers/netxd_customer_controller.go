package controllers

import (
	"context"

	"github.com/balamh/netxd_dal/netxd_dal_interfaces"
	"github.com/balamh/netxd_dal/netxd_dal_models"
	pro "github.com/balamh/project1/netxd_customer"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService netxd_dal_interfaces.INetxdCustomers
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	dbInsert := &netxd_dal_models.NetxdCustomer{
		CustomerId: req.CustomerId,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		BankId:     req.BankId,
		Balance:    float64(req.Balance),
		IsActive:   req.IsActive,
	}
	result, err := CustomerService.CreateCustomer(dbInsert)
	if err != nil {
		return nil, err
	} else {
		Response := &pro.CustomerResponse{
			CustomerId: result.CustromerId,
		}
		return Response, nil
	}
}
