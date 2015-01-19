package events

import (
	"fmt"

	"bytes"

	. "bitbucket.org/abdullin/proto/back/shared"
)

func i(c string, eventId Id) *Info {
	return &Info{c, eventId}
}

type ProductCreated struct {
	EventId   Id        `json:"eventId"`
	ProductId ProductId `json:"productId"`
	Name      string    `json:"name"`
	SkuId     string    `json:"sku"`
}

func (e *ProductCreated) Meta() *Info { return i("ProductCreated", e.EventId) }

func (e *ProductCreated) String() string {
	return fmt.Sprintf("Product '%s' created with sku '%s' as '%s'", e.Name, e.SkuId, e.ProductId)
}

func NewProductCreated(event Id, product ProductId, name string, sku string) *ProductCreated {
	return &ProductCreated{event, product, name, sku}
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

func (e *ItemAdded) String() string {
	return fmt.Sprintf("Added %v of product %s to location %s", e.Quantity, e.ProductId, e.LocationId)
}

type LocationId struct{ Id }

func NewLocationId() LocationId {
	return LocationId{NewId()}
}

type ProductId struct{ Id }

func NewProductId() ProductId {
	return ProductId{NewId()}
}

type LocationCreated struct {
	EventId    Id         `json:"eventId"`
	LocationId LocationId `json:"locationId"`
	Name       string     `json:"name"`
}

func (e *LocationCreated) Meta() *Info { return i("LocationCreated", e.EventId) }
func (e *LocationCreated) String() string {
	return fmt.Sprintf("Location '%s' created as '%s'", e.Name, e.LocationId)
}

func NewLocationCreated(event Id, loc LocationId, name string) *LocationCreated {
	return &LocationCreated{event, loc, name}
}

type VirtualGroupCreated struct {
	EventId Id            `json:"eventId"`
	GroupId ProductId     `json:"recipeId"`
	Name    string        `json:"name"`
	Items   []ProductLine `json:"items"`
}

type ProductLine struct {
	ProductId ProductId `json:"productId"`
	Quantity  int       `json:"quantity"`
}

func NewVirtualGroupCreated(
	event Id,
	recipe ProductId,
	name string,
	items []ProductLine,
) *VirtualGroupCreated {
	return &VirtualGroupCreated{event, recipe, name, items}
}

func (e *VirtualGroupCreated) Meta() *Info { return i("VirtualGroupCreated", e.EventId) }

func (e *VirtualGroupCreated) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Created group '%s' as '%s' with:\n", e.Name, e.GroupId))
	for _, l := range e.Items {
		buffer.WriteString(fmt.Sprintf("    %v of %s\n", l.Quantity, l.ProductId))
	}
	return buffer.String()

}
