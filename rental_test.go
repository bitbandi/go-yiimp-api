package yiimp

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRental(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			}`

	expectedItem := Rental{
	}

	mux.HandleFunc("/api/rental", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "APIKEY", r.URL.Query().Get("key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewYiimpClient(httpClient, "http://dummy.com/", "FAKEKEY", "")
	wallet, err := mposClient.GetRental("APIKEY")

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, wallet)
}
