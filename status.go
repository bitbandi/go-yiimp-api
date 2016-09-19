package yiimp

import (
	"encoding/json"
	"strconv"
	"strings"
)

type PoolStatus map[string]Algo

type Algo struct {
	Name            string `json:"name"`
	Port            uint16 `json:"port"`
	Coins           uint16 `json:"coins"`
	Fees            float32 `json:"fees"`
	Hashrate        uint64 `json:"hashrate,string"`
	Workers         uint16 `json:"workers"`
	EstimateCurrent float64 `json:"estimate_current,string"`
	EstimateLast24h float64 `json:"estimate_last24h,string"`
	ActualLast24h   float64 `json:"actual_last24h,string"`
	RentalCurrent   float64 `json:"rental_current,string"`
	LastBlock       uint32 `json:"lastbloc"`
	TimeSinceLast   uint32 `json:"timesincelast"`
}

func (a *Algo) UnmarshalJSON(data []byte) error {
	type Alias Algo
	aux := &struct {
		ActualLast24h string `json:"actual_last24h"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.ActualLast24h) == 0 {
		a.ActualLast24h = 0
		return nil
	}
	val, err := strconv.ParseFloat(strings.Trim(aux.ActualLast24h, "\""), 64)
	if err != nil {
		return err
	}
	a.ActualLast24h = val
	return nil
}

func (client *YiimpClient) GetStatus() (PoolStatus, error) {
	poolstatus := PoolStatus{}
	_, err := client.sling.New().Get("status").ReceiveSuccess(&poolstatus)
	if err != nil {
		return poolstatus, err
	}

	return poolstatus, err
}
