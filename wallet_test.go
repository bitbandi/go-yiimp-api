package yiimp

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetWallet(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			   "currency":"BTC",
			   "unsold":0,
			   "balance":0.00019171,
			   "unpaid":0.00019171,
			   "paid":7.13839179,
			   "total":7.1385835
			}`

	expectedItem := Wallet{
		Currency:"BTC",
		Unsold:0,
		Balance:0.00019171,
		Unpaid:0.00019171,
		Paid:7.13839179,
		Total:7.1385835}

	mux.HandleFunc("/api/wallet", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "BTCADDRESS", r.URL.Query().Get("address"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewYiimpClient(httpClient, "http://dummy.com/", "FAKEKEY", "")
	wallet, err := mposClient.GetWallet("BTCADDRESS")

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, wallet)
}

func TestGetWalletEx(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			   "currency":"BTC",
			   "unsold":0.0044981783426058,
			   "balance":0.00064099,
			   "unpaid":0.00513917,
			   "paid":0.4005716,
			   "total":0.40571077,
			   "miners":[
			      {
			         "version":"cgminer/1.0.7",
			         "password":"d=0.004",
			         "ID":"",
			         "algo":"x11",
			         "difficulty":0.004,
			         "subscribe":0,
			         "accepted":358763314.0657903,
			         "rejected":153755706.02819583
			      }
			   ]
			}`

	expectedItem := WalletEx{
		Currency:"BTC",
		Unsold:0.0044981783426058,
		Balance:0.00064099,
		Unpaid:0.00513917,
		Paid:0.4005716,
		Total:0.40571077,
		Miners: []Miner{
			Miner{
				Version:"cgminer/1.0.7",
				Password:"d=0.004",
				Id:"",
				Algo:"x11",
				Difficulty:0.004,
				Subscribe:0,
				Accepted:358763314.0657903,
				Rejected:153755706.02819583,
			},
		}}

	mux.HandleFunc("/api/walletex", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "BTCADDRESS", r.URL.Query().Get("address"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewYiimpClient(httpClient, "http://dummy.com/", "FAKEKEY", "")
	wallet, err := mposClient.GetWalletEx("BTCADDRESS")

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, wallet)
}
