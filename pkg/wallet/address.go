package wallet

type AddressGenerator interface {
	Generate(men string, derivation string) (string, error)
}

type AddressStorage interface {
	SaveAddress(wallet *Wallet, address *Address) error
	GetAdresses(wallet *Wallet) ([]Address,error)
}

type Address struct {
	id uint
	code string
	derivation string
}

func (a Address) Id() uint {
	return a.id
}

func (a Address) Derivation() string {
	return a.derivation
}

func (a Address) Code() string {
	return a.code
}

func NewAddressWithFields(id uint, code string, derivation string) *Address {
	return &Address{id: id, code: code, derivation: derivation}
}

func NewAddress(men string, derivation string, generator AddressGenerator) (*Address, error){
	add, err := generator.Generate(men, derivation)
	if err != nil {
		return nil, err
	}

	return &Address{
		code: add,
		derivation: derivation,
	}, nil
}


