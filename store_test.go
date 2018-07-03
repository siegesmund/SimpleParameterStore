package simplestore

import (
		"testing"
		"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	A string `ssm_name:"STORE_STRUCT_TEST_A"`
	B string `ssm_name:"STORE_STRUCT_TEST_B"`
	C string `ssm_name:"STORE_STRUCT_TEST_C" ssm_type:"SecureString"`
}


var store Store

func setupStore() {
	store = Store{Region: "us-east-1"}
}


func TestStoreGetParameter(t *testing.T) {
	param, _ := store.GetParameter("TEST_PLAIN", false)
	assert.Equal(t, "TEST_VALUE_PLAIN", param)
}

func TestStorePutParameter(t *testing.T) {
	version, _ := store.PutParameter("STORE_PUT_TEST", "STORE_PUT_TEST_VALUE", "String","This is a test created parameter", true)
	assert.Equal(t, 1, version)
	param, _ := store.GetParameter("STORE_PUT_TEST", false)
	assert.Equal(t, "STORE_PUT_TEST_VALUE", param)
	store.DeleteParameter("STORE_PUT_TEST")
}


func TestStoreGet(t *testing.T) {
	store.PutParameter("STORE_STRUCT_TEST_A", "STORE_STRUCT_TEST_A_VALUE", "String","This is a test created parameter", true)
	store.PutParameter("STORE_STRUCT_TEST_B", "STORE_STRUCT_TEST_B_VALUE", "String","This is a test created parameter", true)
	store.PutParameter("STORE_STRUCT_TEST_C", "STORE_STRUCT_TEST_C_VALUE", "SecureString","This is a test created parameter", true)

	testStruct := TestStruct{}

	store.Get(&testStruct)

	assert.Equal(t, "STORE_STRUCT_TEST_A_VALUE", testStruct.A)
	assert.Equal(t, "STORE_STRUCT_TEST_B_VALUE", testStruct.B)
	assert.Equal(t, "STORE_STRUCT_TEST_C_VALUE", testStruct.C)


	store.DeleteParameter("STORE_STRUCT_TEST_A")
	store.DeleteParameter("STORE_STRUCT_TEST_B")
	store.DeleteParameter("STORE_STRUCT_TEST_C")
}