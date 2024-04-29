package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func inint() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4,required"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Price < 0 {
		return false, errors.New("the price must be greater than zero")
	}

	if p.Name == "" {
		return false, errors.New("the name is required")
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be ENABLED or DISABLED")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	isValid, err := p.IsValid()
	if err != nil {
		return errors.New("the product must be valid to enable it: " + err.Error())
	}

	if p.Price > 0 && isValid && err == nil {
		p.Status = ENABLED
		return nil
	}

	return errors.New("the price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New("the price must be equal to zero to disable the product")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
