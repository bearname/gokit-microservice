package application

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

var (
	ErrProductNotFound  = errors.New("product not found")
	ErrDuplicateProduct = errors.New("product with such id already exists")
)

type ProductParams interface {
	GetTitle() string
	GetSKU() string
	GetPrice() decimal.Decimal
	GetAvailableCount() int
	GetImageURL() string
	GetImageWidth() int
	GetImageHeight() int
	GetColor() string
	GetMaterial() string
}

type Service interface {
	Create(params ProductParams) (ProductID, error)
	FindByID(id uuid.UUID) (*Product, error)
	Find(page *Page, filters *Filters) ([]*Product, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) *service {
	return &service{repo: repository}
}

func (s *service) Create(params ProductParams) (ProductID, error)  {
	id := s.repo.NextID()
	item := Product{
		ID:             id,
		Title:          params.GetTitle(),
		SKU:            params.GetSKU(),
		Price:          params.GetPrice(),
		AvailableCount: params.GetAvailableCount(),
		Image: &Image{
			URL:    params.GetImageURL(),
			Width:  params.GetImageWidth(),
			Height: params.GetImageHeight(),
		},
		Color:    params.GetColor(),
		Material: params.GetMaterial(),
	}
	err := s.repo.Add(item)
	if err != nil {
		return ProductID{}, err
	}
	return id, nil
}

func (s *service) FindByID(id uuid.UUID) (*Product, error){
	return s.repo.FindByID(ProductID(id))
}

func (s *service) Find(page *Page, filters *Filters) ([]*Product, error) {
	return s.repo.Find(page, filters)
}
