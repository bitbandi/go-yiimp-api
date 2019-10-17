package yiimp

import (
	"encoding/json"
	"errors"
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
	EstimateCurrent float64 `json:"estimate_current"`
	EstimateLast24h float64 `json:"estimate_last24h"`
	ActualLast24h   float64 `json:"actual_last24h"`
	Hashrate24h     float64 `json:"hashrate_last24h"`
	UnitFactor      float64 `json:"mbtc_mh_factor,omitempty"`
	RentalCurrent   float64 `json:"rental_current,omitempty"`
}

func ToFloat64(value interface{}) (float64, error) {
	if value == nil {
		return 0.0, nil
	}
	switch value.(type) {
	case string:
		if len(value.(string)) == 0 {
			return 0.0, nil
		}
		val, err := strconv.ParseFloat(strings.Trim(value.(string), "\""), 64)
		if err != nil {
			return 0.0, err
		}
		return val, nil
	case float64:
		return value.(float64), nil
	case int:
		return float64(value.(int)), nil
	default:
		return 0.0, errors.New("JSON type not understood")
	}
}

func (a *Algo) UnmarshalJSON(data []byte) error {
	type Alias Algo
	aux := &struct {
		EstimateCurrent interface{} `json:"estimate_current"`
		EstimateLast24h interface{} `json:"estimate_last24h"`
		ActualLast24h   interface{} `json:"actual_last24h"`
		Hashrate24h     interface{} `json:"hashrate_last24h"`
		RentalCurrent   interface{} `json:"rental_current,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var err error
	if a.EstimateCurrent, err = ToFloat64(aux.EstimateCurrent); err != nil {
		return err
	}
	if a.EstimateLast24h, err = ToFloat64(aux.EstimateLast24h); err != nil {
		return err
	}
	if a.ActualLast24h, err = ToFloat64(aux.ActualLast24h); err != nil {
		return err
	}
	if a.Hashrate24h, err = ToFloat64(aux.Hashrate24h); err != nil {
		return err
	}
	if a.RentalCurrent, err = ToFloat64(aux.RentalCurrent); err != nil {
		return err
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
