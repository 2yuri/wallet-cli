package address

var CoinMAP map[string]string

func init(){
	CoinMAP = make(map[string]string, 1)

	CoinMAP["ETH"] = "60"
	CoinMAP["TRX"] = "195"
}