package wallet

type TransactionStorage interface {
	CreateTransaction(t *Transaction) error
	UpdateTranscations() error
	FindTransactions(items int) ([]Transaction,error)
}

type TransactionActions interface {
	SendTransaction(t *Transaction) error
	GetFee(t *Transaction) (string, error)
}

type Transaction struct {
	txid string
	amount string
	fee string
	status string
	blockHash string
	blockConfirmatios string
	toAddress string
	currency *Currency
	address *Address
}

func NewTransaction(amount string, toAddress string, currency *Currency, address *Address) *Transaction {
	return &Transaction{amount: amount, status: "pending", toAddress: toAddress, currency: currency, address: address}
}

func (t *Transaction) GetFee(actions TransactionActions) (string, error) {
	return actions.GetFee(t)
}

func (t *Transaction) SendTransaction(actions TransactionActions) error {
	return actions.SendTransaction(t)
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

func (t *Transaction) BlockHash() string {
	return t.blockHash
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

func NewTransactionWithFields(txid string, amount string, fee string, status string, blockHash string, blockConfirmatios string, toAddress string, currency *Currency, address *Address) *Transaction {
	return &Transaction{txid: txid, amount: amount, fee: fee, status: status, blockHash: blockHash, blockConfirmatios: blockConfirmatios, toAddress: toAddress, currency: currency, address: address}
}
