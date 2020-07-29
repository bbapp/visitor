package visitor

import (
	"fmt"
	"github.com/lovetsky/visitor/pkg/account"
)

// Visitor обеспечивает интерфейс покупателя
type Visitor interface {
	VisitFoAccount(l account.Accounter) (err error)
}

type calcAddons struct {
	money int
}

// VisitFoAccount добавляем новый метод
func (a *calcAddons) VisitFoAccount(l account.Accounter) (err error) {
	fmt.Println("Добавили денег на кашелек (новый метод):", a.money)
	return
}

// NewArray visitor factory
func NewVisitor(money int) Visitor {
	return &calcAddons{
		money,
	}
}