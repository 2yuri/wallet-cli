package wallet

type TransactionStorage interface {
	CreateTransaction(t *Transaction) error
	UpdateTranscations() error
	FindTransactions(items int, status string) ([]Transaction,error)
}

type TransactionActions interface {
	SendTransaction(mnemonic string, t *Transaction) (*Transaction, error)
	GetFee(t *Transaction) (string, error)
}

type Transaction struct {
	txid string
	amount string
	fee string
	status string
	blockConfirmatios string
	toAddress string
	currency *Currency
	address *Address
}

func (t *Transaction) SetTxid(txid string) {
	t.txid = txid
}

func (t *Transaction) SetBlockConfirmatios(blockConfirmatios string) {
	t.blockConfirmatios = blockConfirmatios
}

func NewTransaction(amount string, toAddress string, currency *Currency, address *Address) *Transaction {
	return &Transaction{amount: amount, status: "pending", toAddress: toAddress, currency: currency, address: address}
}

func (t *Transaction) GetFee(actions TransactionActions) (string, error) {
	return actions.GetFee(t)
}

func (t *Transaction) SendTransaction(actions TransactionActions, mnemonic string) (*Transaction, error) {
	return actions.SendTransaction(mnemonic, t)
}

func (t *Transaction) Txid() string {
	return t.txid
}

func (t *Transaction) Amount() string {
	return t.amount
}

func (t *Transaction) Fee() string {
	return t.fee
}

func (t *Transaction) Status() string {
	return t.status
}

func (t *Transaction) BlockConfirmatios() string {
	return t.blockConfirmatios
}

func (t *Transaction) ToAddress() string {
	return t.toAddress
}

func (t *Transaction) Currency() *Currency {
	return t.currency
}

func (t *Transaction) Address() *Address {
	return t.address
}

func NewTransactionWithFields(txid string, amount string, fee string, status string, blockConfirmatios string, toAddress string, currency *Currency, address *Address) *Transaction {
	return &Transaction{txid: txid, amount: amount, fee: fee, status: status, blockConfirmatios: blockConfirmatios, toAddress: toAddress, currency: currency, address: address}
}
