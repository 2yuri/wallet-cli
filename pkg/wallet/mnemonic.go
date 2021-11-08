package wallet

type MnemonicGenerator interface {
	Generate() (*Mnemonic, error)
}

type Mnemonic struct {
	code string
}

func (m Mnemonic) Code() string {
	return m.code
}

func NewMnemonic(code string) *Mnemonic {
	return &Mnemonic{code: code}
}
