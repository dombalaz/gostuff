package main

import (
	"testing"
)

// Seems like cap is one time faster than Close
// Close uses wgs
func BenchmarkCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchCap()
	}
}

func BenchmarkClose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchClose()
	}
}
