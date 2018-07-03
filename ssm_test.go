package simplestore

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func setup() {
	putParameter("TEST_SECURE", "TEST_VALUE_SECURE", "SecureString", "This is a test created parameter", "us-east-1", false)
	putParameter("TEST_PLAIN", "TEST_VALUE_PLAIN", "String", "This is a test created parameter", "us-east-1", false)
	putParameter("TEST_DELETE", "TEST_DELETE_PLAIN", "String", "This is a test created parameter", "us-east-1", false)
}

func teardown() {
	deleteParameter("TEST_SECURE", "us-east-1")
	deleteParameter("TEST_PLAIN", "us-east-1")
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}


func testParameterValues(t *testing.T, name, value string, withDecryption bool) ([]*ssm.Parameter, error) {

	params, err := getParameters([]*string{&name}, "us-east-1", withDecryption)

	assert.Nil(t, err, fmt.Sprintf("%s returned an error", name))

	assert.Equal(t, 1, len(params))
	assert.Equal(t, value, *params[0].Value)

	return params, err
}


func TestGetParameters(t *testing.T) {
	testParameterValues(t, "TEST_SECURE", "TEST_VALUE_SECURE", true)
	testParameterValues(t, "TEST_PLAIN", "TEST_VALUE_PLAIN", false)
}

func TestDeleteParameters(t *testing.T) {

	name := "TEST_DELETE"
	// Test that the parameter exists
	params, err := getParameters([]*string{&name}, "us-east-1", false)
	assert.Nil(t, err, fmt.Sprintf("%s returned an error", name))
	assert.Equal(t, 1, len(params))

	_, err = deleteParameter(name, "us-east-1")
	assert.Nil(t, err)

	// Test that the parameter has been deleted
	params, err = getParameters([]*string{&name}, "us-east-1", false)
	assert.Nil(t, err, fmt.Sprintf("%s returned an error", name))
	assert.Equal(t, 0, len(params))
}