package main

import (
	"testing"
)

func TestResponse(t *testing.T) {
	actualRes, _ := handleRequest()
	expectedRes := "Hello λ!"
	if actualRes != expectedRes {
		t.Errorf("Response Missmatch: Got %s, want %s", actualRes, expectedRes)
	}
}

func TestError(t *testing.T) {
	_, actualErr := handleRequest()
	expectedErr := error(nil)
	if actualErr != expectedErr {
		t.Errorf("Error Missmatch: Got %v, want %v", actualErr, expectedErr)
	}
}
