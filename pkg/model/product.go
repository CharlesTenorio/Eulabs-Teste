package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/eulabs/back-end/api-product/internal/config/logger"
	"github.com/twinj/uuid"
)

type Product struct {
	ID         int64     `json:"-"`
	ExternalID string    `json:"id"`
	Name       string    `json:"name"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

func (p *Product) String() string {
	data, err := json.Marshal(p)
	if err != nil {
		logger.Error("Error to marshal product: %s", err)
		return ""
	}
	return string(data)
}

type ProductList struct {
	List []*Product `json:"list"`
}

func NewProduct(prod_request *Product) (*Product, error) {
	Product := &Product{
		ID:         prod_request.ID,
		ExternalID: uuid.NewV4().String(),
		Name:       prod_request.Name,
		Quantity:   prod_request.Quantity,
		Price:      prod_request.Price,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := Product.invalid(); err != nil {
		return nil, err
	}
	return Product, nil

}

func (p *Product) invalid() error {
	if p.Quantity <= 0 {
		return fmt.Errorf("operator id is required")
	}
	if p.Name == "" {
		return fmt.Errorf("status customer service id is required")
	}

	return nil
}
