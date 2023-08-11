package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Can't find the product")
	ErrCantDecodeProducts = errors.New("Can't find the products")
	ErrUserIdIsNotValid   = errors.New("This user is not valid")
	ErrCantUpdateUser     = errors.New("Cannot add this product to the cart")
	ErrCantRemoveItem     = errors.New("Cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("Was unnable to get the item from the cart")
	ErrCantBuyItem        = errors.New("Cannot update the purchase")
)

func AddToCart() {

}

func RemoveFromCart() {

}

func BuyFromCart() {

}

func Buyer() {

}
