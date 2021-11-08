package wallet_cli

type AddressGenerator interface {
	Generate(men string, derivation string, network string) (*Address, error)
}

type Address struct {
	code string
}

func (a Address) Code() string {
	return a.code
}

func NewAddress(code string) *Address {
	return &Address{code: code}
}
