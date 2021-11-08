package wallet

type BalanceInfo interface {
	GetBalance(a *Address) (*Balance, error)
}

type Balance struct {
	confimated string
	unconfirmed string
}

func (b Balance) Confimated() string {
	return b.confimated
}

func (b Balance) Unconfirmed() string {
	return b.unconfirmed
}

func NewBalance(confimated string, unconfirmed string) *Balance {
	return &Balance{confimated: confimated, unconfirmed: unconfirmed}
}
