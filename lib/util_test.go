package lib

import (
	"strconv"
	"testing"
)

const knownData = "test data"

// Calculated using ISO table
const knownHash = 10232006911339297906

func TestHashWorks(t *testing.T) {
	if HashCRC64(knownData) != knownHash {
		t.Fatalf("Invalid hash returned")
	}
}

// Test for correctness
func TestHashIsConsistent(t *testing.T) {
	ret := HashCRC64(knownData)
	if ret != knownHash {
		t.Errorf("Invalid hash returned")
	}
}

// Test for consistency when function is used for other data
func TestHashIsConsistentAcrossMultipleRuns(t *testing.T) {
	for i := 0; i < 50000; i++ {
		HashCRC64(strconv.Itoa(i))
	}

	ret := HashCRC64(knownData)
	if ret != knownHash {
		t.Errorf("Invalid hash returned")
	}
}

func TestHasAuthPrefix(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		scheme   string
		expected bool
	}{
		{name: "basic with space", token: "Basic Zm9vOmJhcg==", scheme: "Basic", expected: true},
		{name: "basic lowercase", token: "basic Zm9vOmJhcg==", scheme: "Basic", expected: true},
		{name: "basic with tab", token: "Basic\tZm9vOmJhcg==", scheme: "Basic", expected: true},
		{name: "basic only scheme", token: "Basic ", scheme: "Basic", expected: true},
		{name: "missing space", token: "BasicZm9v", scheme: "Basic", expected: false},
		{name: "different scheme", token: "Bot foo", scheme: "Basic", expected: false},
		{name: "bearer", token: "Bearer token", scheme: "Bearer", expected: true},
	}

	for _, tc := range tests {
		if result := HasAuthPrefix(tc.token, tc.scheme); result != tc.expected {
			t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
		}
	}
}
