package currencyapi_test

import (
	"testing"

	"maxklammer.com/go/cryptomasters/currencyapi"
)

func TestGetRate(t *testing.T) {
	_, err := currencyapi.GetRate("")
	if err == nil {
		t.Errorf("Empty currencies should return error")
	}
}
