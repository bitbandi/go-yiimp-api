package yiimp

type Rental struct {
	Balance     float64 `json:"balance"`
	Unconfirmed float64 `json:"unconfirmed"`
	Jobs        RentalJob `json:"jobs"`
}

type RentalJob struct {
	JobId      uint32 `json:"jobid,string"`
	Algo       string `json:"algo"`
	Price      float64 `json:"price,string"`
	Hashrate   uint64 `json:"hashrate,string"`
	Server     string `json:"server"`
	Port       uint16 `json:"port,string"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Started    uint8 `json:"started"`
	Active     uint8 `json:"active"`
	Accepted   float64 `json:"accepted,string"`
	Rejected   float64 `json:"rejected,string"`
	Difficulty float32 `json:"diff,string"`
}

type yiimpRentalRequest struct {
	Key string `url:"key"`
}

func (client *YiimpClient) GetRental(key string) (Rental, error) {
	rental := Rental{}
	req := &yiimpRentalRequest{Key: key}
	_, err := client.sling.New().Get("rental").QueryStruct(req).ReceiveSuccess(&rental)
	if err != nil {
		return rental, err
	}

	return rental, err
}
