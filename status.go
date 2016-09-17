package yiimp

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

func (client *YiimpClient) GetStatus() (PoolStatus, error) {
	poolstatus := PoolStatus{}
	_, err := client.sling.New().Get("status").ReceiveSuccess(&poolstatus)
	if err != nil {
		return poolstatus, err
	}

	return poolstatus, err
}
