package main

import (
    "context"
    pb "github.com/arturfil/m_commons/api"
)

type PaymentService interface {
    CreatePayment(context.Context, *pb.Order) (string,error)
} 
