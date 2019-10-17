package yiimp

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetStatus(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			   "test1": {
			      "name": "test1",
			      "port": 4234,
			      "coins": 3,
			      "fees": 2,
			      "hashrate": 365873,
			      "workers": 26,
			      "estimate_current": "0.01343157",
			      "estimate_last24h": "0.01821587",
			      "actual_last24h": "17.73317",
			      "hashrate_last24h": 29817174797.385,
			      "mbtc_mh_factor": 1000,
			      "rental_current": "13.97979333"
			   },
			   "test2": {
			      "name": "test2",
			      "port": 5766,
			      "coins": 3,
			      "fees": 2,
			      "hashrate": 11304269627,
			      "workers": 13,
			      "estimate_current": "0.00181736",
			      "estimate_last24h": "0.00164080",
			      "actual_last24h": "",
			      "hashrate_last24h": 10668316988.396,
			      "rental_current": "1.95910447"
			   },
			   "test3": {
			      "name": "test3",
			      "port": 6576,
			      "coins": 3,
			      "fees": 2,
			      "hashrate": 2863888,
			      "workers": 4,
			      "estimate_current": 0.00000339,
			      "estimate_last24h": 0.00000433,
			      "actual_last24h": 0.0037,
			      "hashrate_last24h": 1876221955.0729,
			      "rental_current": 1.36776496
			   },
			   "test4": {
			      "name": "test4",
			      "port": 8453,
			      "coins": 3,
			      "fees": 2,
			      "hashrate": 2863888,
			      "workers": 4,
			      "estimate_current": 0.00000339,
			      "estimate_last24h": 0.00000433,
			      "actual_last24h": 0.0037,
			      "hashrate_last24h": 1876221955.0729
			   }
			}`

	expectedItem := PoolStatus{
		"test1": Algo{
			Name: "test1",
			Port: 4234,
			Coins: 3,
			Fees: 2,
			Hashrate: 365873,
			Workers: 26,
			EstimateCurrent: 0.01343157,
			EstimateLast24h: 0.01821587,
			ActualLast24h: 17.73317,
			Hashrate24h: 29817174797.385,
			UnitFactor: 1000,
			RentalCurrent: 13.97979333,
		},
		"test2": Algo{
			Name: "test2",
			Port: 5766,
			Coins: 3,
			Fees: 2,
			Hashrate: 11304269627,
			Workers: 13,
			EstimateCurrent: 0.00181736,
			EstimateLast24h: 0.00164080,
			ActualLast24h: 0,
			Hashrate24h: 10668316988.396,
			UnitFactor: 0,
			RentalCurrent: 1.95910447,
		},
		"test3": Algo{
			Name: "test3",
			Port: 6576,
			Coins: 3,
			Fees: 2,
			Hashrate: 2863888,
			Workers: 4,
			EstimateCurrent: 0.00000339,
			EstimateLast24h: 0.00000433,
			ActualLast24h: 0.0037,
			Hashrate24h: 1876221955.0729,
			UnitFactor: 0,
			RentalCurrent: 1.36776496,
		},
		"test4": Algo{
			Name: "test4",
			Port: 8453,
			Coins: 3,
			Fees: 2,
			Hashrate: 2863888,
			Workers: 4,
			EstimateCurrent: 0.00000339,
			EstimateLast24h: 0.00000433,
			ActualLast24h: 0.0037,
			Hashrate24h: 1876221955.0729,
			UnitFactor: 0,
			RentalCurrent: 0,
		},
	}

	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	yiimpClient := NewYiimpClient(httpClient, "http://dummy.com/", "FAKEKEY", "")
	status, err := yiimpClient.GetStatus()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, status)
}