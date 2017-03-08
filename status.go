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
	Hashrate        float64 `json:"hashrate"`
	Workers         uint16 `json:"workers"`
	EstimateCurrent float64 `json:"estimate_current,string"`
	EstimateLast24h float64 `json:"estimate_last24h,string"`
	ActualLast24h   float64 `json:"actual_last24h,string"`
	Hashrate24h     float64 `json:"hashrate_last24h"`
	RentalCurrent   float64 `json:"rental_current,string"`
}

func (a *Algo) UnmarshalJSON(data []byte) error {
	type Alias Algo
	aux := &struct {
		ActualLast24h interface{} `json:"actual_last24h"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch aux.ActualLast24h.(type) {
	case string:
		if len(aux.ActualLast24h.(string)) == 0 {
			a.ActualLast24h = 0
			return nil
		}
		val, err := strconv.ParseFloat(strings.Trim(aux.ActualLast24h.(string), "\""), 64)
		if err != nil {
			return err
		}
		a.ActualLast24h = val
	case float64:
		a.ActualLast24h = aux.ActualLast24h.(float64)
	case int:
		a.ActualLast24h = float64(aux.ActualLast24h.(int))
	default:
		panic("JSON type not understood")
	}
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
