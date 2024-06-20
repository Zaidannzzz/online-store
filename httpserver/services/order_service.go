package services

import (
	"errors"
	"online-store/httpserver/repositories"
)

type OrderService interface {
	Checkout(userID uint) error
}

type orderService struct {
	cartRepo  repositories.CartRepository
	orderRepo repositories.OrderRepository
}

func NewOrderService(cartRepo repositories.CartRepository, orderRepo repositories.OrderRepository) OrderService {
	return &orderService{cartRepo: cartRepo, orderRepo: orderRepo}
}

func (s *orderService) Checkout(userID uint) error {
	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	if len(cart.CartItems) == 0 {
		return errors.New("cart is empty")
	}

	err = s.orderRepo.CreateOrder(userID, cart.CartItems, cart.Total)
	if err != nil {
		return err
	}

	return s.cartRepo.ClearCart(userID)
}
