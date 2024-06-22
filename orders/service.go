package main

import (
	"context"
	"log"

	common "github.com/arturfil/m_commons"
	pb "github.com/arturfil/m_commons/api"
)

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	items, err := s.ValidateOrder(ctx, p)
	if err != nil {
		return nil, err
	}

	o := &pb.Order{
		ID:         "42",
		CustomerID: p.CustomerID,
		Status:     "pending",
		Items:      items,
	}

	return o, nil
}

func (s *service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.Item, error) {
	if len(p.Items) == 0 {
		return nil, common.ErrNoItems
	}

	mergedItems := mergeItemsQuantities(p.Items)
	log.Print(mergedItems)

	var itemsWithPrice []*pb.Item
	for _, i := range mergedItems {
		itemsWithPrice = append(itemsWithPrice, &pb.Item{
			PriceID:  "price_1PU0koIAKhlfjLC8hQpbqfXk",
			ID:       i.ID,
			Quantity: i.Quantity,
		})
	}

	return itemsWithPrice, nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	merged := make([]*pb.ItemsWithQuantity, 0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.ID == item.ID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}

		if !found {
			merged = append(merged, item)
		}
	}

	return merged
}
