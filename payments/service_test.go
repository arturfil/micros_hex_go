package main

import (
	"context"
	"testing"

	"github.com/arturfil/m_commons/api"
	"github.com/arturfil/m_payments/processor/inmem"
)


func TestService(t *testing.T) {
    processor := inmem.NewInmen()
    svc := NewService(processor)

    t.Run("should create a payment link", func(t *testing.T) {
        link, err := svc.CreatePayment(context.Background(), &api.Order{})
        if err != nil {
            t.Errorf("CreatePayment() error = %v, want nil", err)
        }

        if link == "" {
            t.Error("CreatePayment() link is empty")
        }
    })
}
