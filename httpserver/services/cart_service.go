package services

import (
	"online-store/httpserver/dto"
	"online-store/httpserver/models"
	"online-store/httpserver/repositories"
)

type CartService interface {
	AddToCart(ID uint, cartItemDTO dto.CartItemDTO) error
	GetCartByUserID(userID uint) (*dto.CartDTO, error)
	RemoveFromCart(cartItemID uint) error
	ClearCart(userID uint) error
}

type cartService struct {
	cartRepo repositories.CartRepository
}

func NewCartService(cartRepo repositories.CartRepository) CartService {
	return &cartService{cartRepo: cartRepo}
}

func (s *cartService) AddToCart(ID uint, AddToCartDto dto.CartItemDTO) error {
	cartItem := models.CartItem{
		ProductID: AddToCartDto.ProductID,
		Quantity:  AddToCartDto.Quantity,
		Price:     AddToCartDto.Price,
	}

	cart, err := s.cartRepo.GetCartByUserID(ID)
	if err != nil {
		return err
	}

	cart.Total += cartItem.Price * float64(cartItem.Quantity)
	cart.CartItems = append(cart.CartItems, cartItem)

	return s.cartRepo.AddToCart(&cartItem)
}

func (s *cartService) GetCartByUserID(userID uint) (*dto.CartDTO, error) {
	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	cartDTO := dto.CartDTO{
		ID:        cart.ID,
		UserID:    cart.UserID,
		Total:     cart.Total,
		CartItems: []dto.CartItemDTO{},
	}

	for _, item := range cart.CartItems {
		cartDTO.CartItems = append(cartDTO.CartItems, dto.CartItemDTO{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	return &cartDTO, nil
}

func (s *cartService) RemoveFromCart(cartItemID uint) error {
	return s.cartRepo.RemoveFromCart(cartItemID)
}

func (s *cartService) ClearCart(userID uint) error {
	return s.cartRepo.ClearCart(userID)
}
