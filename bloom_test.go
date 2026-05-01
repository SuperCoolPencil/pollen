package bloom

import (
	"fmt"
	"testing"
)

func TestBloomFilter_Basic(t *testing.T) {
	f := New(1000, 7)

	data1 := []byte("hello")
	data2 := []byte("world")
	data3 := []byte("pollen")

	f.Add(data1)
	f.Add(data2)

	if !f.Check(data1) {
		t.Errorf("Expected 'hello' to be present")
	}
	if !f.Check(data2) {
		t.Errorf("Expected 'world' to be present")
	}
	if f.Check(data3) {
		t.Errorf("Expected 'pollen' to be absent (highly likely)")
	}
}

func TestBloomFilter_Reset(t *testing.T) {
	f := New(1000, 7)
	data := []byte("test")

	f.Add(data)
	if !f.Check(data) {
		t.Errorf("Expected 'test' to be present before reset")
	}

	f.Reset()
	if f.Check(data) {
		t.Errorf("Expected 'test' to be absent after reset")
	}
}

func TestBloomFilter_Estimates(t *testing.T) {
	n := uint64(1000)
	p := 0.01 // 1% false positive rate
	f := NewWithEstimates(n, p)

	// Verify m and k are calculated (non-zero)
	if f.m == 0 {
		t.Errorf("Expected non-zero m from estimates")
	}
	if f.k == 0 {
		t.Errorf("Expected non-zero k from estimates")
	}

	// Add 1000 items
	for i := 0; i < int(n); i++ {
		f.Add([]byte(fmt.Sprintf("item-%d", i)))
	}

	// Check they all exist
	for i := 0; i < int(n); i++ {
		if !f.Check([]byte(fmt.Sprintf("item-%d", i))) {
			t.Errorf("Expected item-%d to be present", i)
		}
	}
}

func TestBloomFilter_FalsePositiveRate(t *testing.T) {
	n := uint64(1000)
	p := 0.05 // 5% false positive rate
	f := NewWithEstimates(n, p)

	// Add n items
	for i := uint64(0); i < n; i++ {
		f.Add([]byte(fmt.Sprintf("element-%d", i)))
	}

	// Test for items not in the filter
	falsePositives := 0
	testSize := 10000
	for i := 0; i < testSize; i++ {
		if f.Check([]byte(fmt.Sprintf("other-%d", i))) {
			falsePositives++
		}
	}

	actualP := float64(falsePositives) / float64(testSize)
	t.Logf("False Positives: %d/%d (Actual P: %f, Target P: %f)", falsePositives, testSize, actualP, p)

	// Allowing some margin for error as it's probabilistic
	if actualP > p*2 {
		t.Errorf("False positive rate too high: %f (target %f)", actualP, p)
	}
}

func TestBloomFilter_Empty(t *testing.T) {
	f := New(1000, 7)
	if f.Check([]byte("")) {
		t.Errorf("Expected empty string to be absent initially")
	}

	f.Add([]byte(""))
	if !f.Check([]byte("")) {
		t.Errorf("Expected empty string to be present after adding")
	}
}

func TestBloomFilter_LargeData(t *testing.T) {
	f := New(1000, 7)
	largeData := make([]byte, 1024*1024) // 1MB
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	f.Add(largeData)
	if !f.Check(largeData) {
		t.Errorf("Expected large data to be present")
	}
}
