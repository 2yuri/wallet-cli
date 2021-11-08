package wallet_cli

type Wallet struct {
	address *Address
	derivation string
	network string
}

func (w Wallet) Address() *Address {
	return w.address
}

func (w Wallet) Derivation() string {
	return w.derivation
}

func (w Wallet) Network() string {
	return w.network
}

func NewWallet(men string, derivation string, network string, generator AddressGenerator) (*Wallet, error){

	add, err := generator.Generate(men, derivation, network)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		address: add,
		derivation: derivation,
		network: network,
	}, nil
}
