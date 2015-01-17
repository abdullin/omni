package events

import (
	. "bitbucket.org/abdullin/proto/back/shared"
)

func i(c string, eventId Id) *Info {
	return &Info{c, eventId}
}

type ProductCreated struct {
	EventId   Id        `json:"eventId"`
	ProductId ProductId `json:"productId"`
	SkuId     string    `json:"sku"`
}

func (e *ProductCreated) Meta() *Info { return i("ProductCreated", e.EventId) }

func NewProductCreated(event Id, product ProductId, sku string) *ProductCreated {
	return &ProductCreated{event, product, sku}
}

type ItemAdded struct {
	EventId    Id         `json:"eventId"`
	ProductId  ProductId  `json:"productId"`
	LocationId LocationId `json:"locationId"`
	Quantity   int        `json:"quantity"`
}

func (e *ItemAdded) Meta() *Info { return i("ItemAdded", e.EventId) }

func NewItemAdded(event Id, product ProductId, location LocationId, quantity int) *ItemAdded {
	return &ItemAdded{event, product, location, quantity}
}

type LocationId Id

func NewLocationId() LocationId {
	return LocationId(NewId())
}

type ProductId Id

func NewProductId() ProductId {
	return ProductId(NewId())
}

type LocationCreated struct {
	EventId    Id         `json:"eventId"`
	LocationId LocationId `json:"locationId"`
	Name       string     `json:"name"`
}

func (e *LocationCreated) Meta() *Info { return i("LocationCreated", e.EventId) }

func NewLocationCreated(event Id, loc LocationId, name string) *LocationCreated {
	return &LocationCreated{event, loc, name}
}

type VirtualGroupCreated struct {
	EventId Id          `json:"eventId"`
	GroupId ProductId   `json:"recipeId"`
	Name    string      `json:"name"`
	Items   ProductList `json:"items"`
}

type ProductList map[ProductId]int

func NewVirtualGroupCreated(
	event Id,
	recipe ProductId,
	name string,
	items ProductList,
) *VirtualGroupCreated {
	return &VirtualGroupCreated{event, recipe, name, items}
}

func (e *VirtualGroupCreated) Meta() *Info { return i("VirtualGroupCreated", e.EventId) }
