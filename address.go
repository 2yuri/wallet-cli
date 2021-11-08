package wallet_cli

type AddressGenerator interface {
	Generate(men string, derivation string, network string) (string, error)
}

type Address struct {
	code string
	derivation string
	network string
}

func (a Address) Code() string {
	return a.code
}

func NewAddress(men string, derivation string, network string, generator AddressGenerator) (*Address, error){
	add, err := generator.Generate(men, derivation, network)
	if err != nil {
		return nil, err
	}

	return &Address{
		code: add,
		derivation: derivation,
		network: network,
	}, nil
}
