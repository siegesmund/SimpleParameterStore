package simplestore

import "testing"

func TestMain(m *testing.M) {
	setupSSM()
	setupStore()
	m.Run()
	teardownSSM()
}


