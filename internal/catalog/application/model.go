package application

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductID uuid.UUID

func (u ProductID) String() string {
	return uuid.UUID(u).String()
}

type Page struct {
	Size   int
	Number int
}

type DecimalRangeFilter struct {
	Min *decimal.Decimal
	Max *decimal.Decimal
}

type StringOrFilter []string

type Filters struct {
	Price    DecimalRangeFilter
	Color    *[]string
	Material *[]string
}

type Image struct {
	URL    string
	Width  int
	Height int
}

type Product struct {
	ID             ProductID
	SKU            string
	Title          string
	Price          decimal.Decimal
	AvailableCount int
	Image          *Image
	Color          string
	Material       string
}

type Repository interface {
	NextID() ProductID
	FindByID(id ProductID) (*Product, error)
	Find(spec *Page, filters *Filters) ([]*Product, error)
	Add(item Product) error
}
