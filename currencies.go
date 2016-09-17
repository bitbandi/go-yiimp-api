package yiimp

type Currencies map[string]Currency

type Currency struct {
	Name          string `json:"name"`
	Symbol        string `json:"symbol,omitempty"`
	Port          uint16 `json:"port"`
	Algo          string `json:"algo"`
	Height        uint32 `json:"height"`
	Workers       uint16 `json:"workers"`
	Shares        uint64 `json:"shares"`
	Hashrate      uint64 `json:"hashrate"`
	LastBlock     uint32 `json:"lastblock"`
	TimeSinceLast uint32 `json:"timesincelast"`
}

func (client *YiimpClient) GetCurrencies() (Currencies, error) {
	currencies := Currencies{}
	_, err := client.sling.New().Get("currencies").ReceiveSuccess(&currencies)
	if err != nil {
		return currencies, err
	}

	return currencies, err
}
