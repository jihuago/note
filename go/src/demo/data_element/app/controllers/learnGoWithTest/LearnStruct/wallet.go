package LearnStruct

import (
	"errors"
	"fmt"
)

// Go允许从现有的类型创建新的类型（类型别名）。
// 创建新的类型，可以让类型描述性
type Bitcoin int

type Stringer interface {
	String() string
}

// 类型别名有一个特性，可以对它们声明方法。当你希望在现有类型之上添加一些领域内特定功能时，这将非常有用
func (b Bitcoin) String() string  {
	return fmt.Sprintf("%d BTC", b)
}

// 声明一个Wallet结构体，利用它存放Bitcoin
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin  {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error  {
	if amount > w.balance {
		return errors.New("oh no")
	}

	w.balance -= amount
	return nil

}

