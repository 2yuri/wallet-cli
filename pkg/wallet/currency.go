package wallet

type CurrencyStorage interface {
	GetCurrency(net, sym string) (*Currency, error)
}

type Currency struct {
	id uint
	symbol string
	network string
	tokenType string
	uri string
	contractAddress string
}

func (c Currency) Symbol() string {
	return c.symbol
}

func (c Currency) Network() string {
	return c.network
}

func (c Currency) TokenType() string {
	return c.tokenType
}

func (c Currency) Uri() string {
	return c.uri
}

func (c Currency) ContractAddress() string {
	return c.contractAddress
}
func (c *Currency) Id() uint {
	return c.id
}

func NewCurrency(id uint, symbol string, network string, tokenType string, uri string, contractAddress string) *Currency {
	return &Currency{id:id,symbol: symbol, network: network, tokenType: tokenType, uri: uri, contractAddress: contractAddress}
}