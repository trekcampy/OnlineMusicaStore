package dblayer

import (
	"models"
)

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(int) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersById(int) ([]models.Order, error)
	AddOrder(models.Order) error
	GetCreditCardCID(int) (string, error)
	SaveCreditCardForCustomer(int, string) error
}
