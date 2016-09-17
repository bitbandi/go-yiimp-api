package yiimp

type Wallet struct {
	Currency string `json:"currency"`
	Unsold   float64 `json:"unsold"`
	Balance  float64 `json:"balance"`
	Unpaid   float64 `json:"unpaid"`
	Paid     float64 `json:"paid"`
	Total    float64 `json:"total"`
}

type WalletEx struct {
	Currency string `json:"currency"`
	Unsold   float64 `json:"unsold"`
	Balance  float64 `json:"balance"`
	Unpaid   float64 `json:"unpaid"`
	Paid     float64 `json:"paid"`
	Total    float64 `json:"total"`
	Miners   []Miner `json:"miners"`
}

type Miner struct {
	Version    string `json:"version"`
	Password   string `json:"password"`
	Id         string `json:"ID"`
	Algo       string `json:"algo"`
	Difficulty float32 `json:"difficulty"`
	Subscribe  uint8 `json:"subscribe"`
	Accepted   float64 `json:"accepted"`
	Rejected   float64 `json:"rejected"`
}

type yiimpWalletRequest struct {
	Address string `url:"address"`
}

func (client *YiimpClient) GetWallet(address string) (Wallet, error) {
	wallet := Wallet{}
	req := &yiimpWalletRequest{Address: address}
	_, err := client.sling.New().Get("wallet").QueryStruct(req).ReceiveSuccess(&wallet)
	if err != nil {
		return wallet, err
	}

	return wallet, err
}

func (client *YiimpClient) GetWalletEx(address string) (WalletEx, error) {
	wallet := WalletEx{}
	req := &yiimpWalletRequest{Address: address}
	_, err := client.sling.New().Get("walletex").QueryStruct(req).ReceiveSuccess(&wallet)
	if err != nil {
		return wallet, err
	}

	return wallet, err
}
