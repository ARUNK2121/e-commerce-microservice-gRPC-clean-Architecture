package interfaces

type CartRepository interface {
	FindCart(int) (int, error)
	CreateCart(int) (int, error)
	CheckIfAlreadyExists(int, int) (bool, error)
	AddToLineItems(int, int) error
	GetProductIDsFromCart(cart int) ([]int, error)
}
