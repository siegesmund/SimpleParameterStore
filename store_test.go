package simplestore

import (
		"testing"
		"github.com/stretchr/testify/assert"
)

var store Store

func setupStore() {
	store = Store{Region: "us-east-1"}
}


func TestStoreGetParameter(t *testing.T) {
	param, _ := store.GetParameter("TEST_PLAIN", false)
	assert.Equal(t, "TEST_VALUE_PLAIN", param)
}

func TestStorePutParameter(t *testing.T) {
	version, _ := store.PutParameter("STORE_PUT_TEST", "STORE_PUT_TEST_VALUE", "String","", true)
	assert.Equal(t, 1, version)
	param, _ := store.GetParameter("STORE_PUT_TEST", false)
	assert.Equal(t, "STORE_PUT_TEST_VALUE", param)
	store.DeleteParameter("STORE_PUT_TEST")
}