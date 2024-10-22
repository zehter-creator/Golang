package chapter_1

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"log"
	"time"
)

type item struct {
	ID string
}
type cart struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	lockedAt     time.Time
	Items        []item
	CurrencyCode string
	isLocked     bool
}

func (c *cart) TotalPrice() (*money.Money, error) {
	return nil, nil
}

func (c *cart) Lock() error {
	return nil
}

func (c *cart) Unlock() error {
	return nil
}
func main() {
	newCart := &cart{}
	totalPrice, err := newCart.TotalPrice()
	if err != nil {
		fmt.Println("İmpossible to compute price of the cart: %s", err)
		return
	}
	log.Println("total price", totalPrice.Display())

	err = newCart.Lock()
	if err != nil {
		log.Println("impossible to lock the cart ", err)
	}
}
