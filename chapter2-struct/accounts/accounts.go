package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// Go에서 이건 receiver, 여기서 receiver는 a
// receiver 작성하는데 struct의 첫 글자를 따서 소문자로 지어야한다.
// Go에서는 이게 method랑 동등한 관계, method를 struct 안에 작성하지 않음.
// Go는 보안을 위해 복사본을 만든다, 실제 account의 변경을 원하면 *를 추가
// Go에게 account나 receiver를 복사하지말고 실제 receive에게 를 주라고 하는거다 == pointer receiver
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
// 아무것도 변경하지 않을 것이라서 복사본을 사용하여도 된다.
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Go가 내부적으로 호출하는 method를 사용하는 방법
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \n Has: ", a.Balance())
}
