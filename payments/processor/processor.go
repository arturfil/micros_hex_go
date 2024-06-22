package processor

import pb "github.com/arturfil/m_commons/api"

type PaymentProcessor interface {
    CreatePaymentLink(*pb.Order) (string, error)
}
