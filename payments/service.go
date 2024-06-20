package main

import (
    "context"
    pb "github.com/arturfil/m_commons/api"
)

type service struct {

}

func NewService() *service {
    return &service{}
}

func (s *service) CreatePayment(context.Context, *pb.Order) (string, error) {
    return "something", nil
}
