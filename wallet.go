package yiimp

import (
	"encoding/json"
	"strconv"
	"strings"
)

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

func (a *Miner) UnmarshalJSON(data []byte) error {
	type Alias Miner
	aux := &struct {
		Subscribe json.Number `json:"subscribe"`
		Accepted  json.Number `json:"accepted"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Subscribe) == 0 {
		a.Subscribe = 0
	} else {
		val, err := strconv.ParseUint(strings.Trim(string(aux.Subscribe), "\""), 10, 8)
		if err != nil {
			return err
		}
		a.Subscribe = uint8(val)
	}
	if len(aux.Accepted) == 0 {
		a.Accepted = 0
	} else {
		val, err := strconv.ParseFloat(strings.Trim(string(aux.Accepted), "\""), 64)
		if err != nil {
			return err
		}
		a.Accepted = val
	}
	return nil
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
