package inmem

import pb "github.com/arturfil/m_commons/api"

type Inmem struct {}

func NewInmen() *Inmem {
    return &Inmem{}
}

func (i *Inmem) CreatePaymentLink(*pb.Order) (string, error) {
    return "dummy-link", nil
}
