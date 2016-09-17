package yiimp

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrencies(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			   "test1": {
			      "algo": "scrypt",
			      "port": 3433,
			      "name": "test1",
			      "height": 1790203,
			      "workers": 76,
			      "shares": 8,
			      "hashrate": 319876314,
			      "lastblock": 1790203,
			      "timesincelast": 41
			   },
			   "test2": {
			      "algo": "x13",
			      "port": 3633,
			      "name": "test2",
			      "height": 93010,
			      "workers": 23,
			      "shares": 2,
			      "hashrate": 0,
			      "lastblock": 92566,
			      "timesincelast": 18275
			   }
			}`

	expectedItem := Currencies{
		"test1": Currency{
			Algo:"scrypt",
			Port:3433,
			Name:"test1",
			Height:1790203,
			Workers:76,
			Shares:8,
			Hashrate:319876314,
			LastBlock:1790203,
			TimeSinceLast:41,
		},
		"test2": Currency{
			Algo:"x13",
			Port:3633,
			Name:"test2",
			Height:93010,
			Workers:23,
			Shares:2,
			Hashrate:0,
			LastBlock:92566,
			TimeSinceLast:18275,
		},
	}

	mux.HandleFunc("/api/currencies", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	yiimpClient := NewYiimpClient(httpClient, "http://dummy.com/", "FAKEKEY", "")
	currencies, err := yiimpClient.GetCurrencies()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, currencies)
}
