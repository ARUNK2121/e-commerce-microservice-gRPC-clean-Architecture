package clienterface

import "order/pkg/utils/models"

type CartClient interface {
	GetCart(user_id int) ([]models.GetCart, error)
}
