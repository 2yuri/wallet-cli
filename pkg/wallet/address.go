package wallet

type AddressGenerator interface {
	Generate(men string, derivation string, network string) (string, error)
}

type AddressStorage interface {
	SaveAddress(wallet *Wallet, address *Address) error
	GetAdresses(wallet *Wallet) ([]Address,error)
}

type Address struct {
	code string
	derivation string
	network string
}

func (a Address) Derivation() string {
	return a.derivation
}

func (a Address) Network() string {
	return a.network
}

func NewAddressWithFields(code string, derivation string, network string) *Address {
	return &Address{code: code, derivation: derivation, network: network}
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
