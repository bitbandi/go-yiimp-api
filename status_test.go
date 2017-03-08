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
			RentalCurrent: 1.95910447,

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