package events

import (
	. "bitbucket.org/abdullin/proto/back/shared"
)

func i(c Contract, eventId Id) *Info {
	return &Info{c, eventId}
}

type ProductCreated struct {
	EventId   Id     `json:"eventId"`
	ProductId Id     `json:"productId"`
	SkuId     string `json:"sku"`
}

func (e *ProductCreated) Meta() *Info { return i("ProductCreated", e.EventId) }

type ItemAdded struct {
	EventId   Id  `json:"eventId"`
	ProductId Id  `json:"productId"`
	Quantity  int `json:"quantity"`
	Location  Id  `json:"location"`
}

func (e *ItemAdded) Meta() *Info { return i("ItemAdded", e.EventId) }

type LocationCreated struct {
	EventId    Id     `json:"eventId"`
	LocationId Id     `json:"locationId"`
	Name       string `json:"name"`
}

func (e *LocationCreated) Meta() *Info { return i("LocationCreated", e.EventId) }
