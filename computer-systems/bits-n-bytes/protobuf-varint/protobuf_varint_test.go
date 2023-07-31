package protobufvarint

import (
	"bytes"
	"math"
	"os"
	"testing"

	"cs-primer/tester"
)

func TestEncode(t *testing.T) {
	// 1
	n1, err := os.ReadFile("./1.uint64")
	tester.NilErr(t, err)
	res := encode(1)
	if !bytes.Equal(res, n1) {
		t.Fatalf("\nWant: %#x\n Got: %#x", n1, res)
	}
	// 150
	n150, err := os.ReadFile("./150.uint64")
	tester.NilErr(t, err)
	res = encode(150)
	if !bytes.Equal(res, n150) {
		t.Fatalf("\nWant: %#x\n Got: %#x", n150, res)
	}
	// Max uint
	nMax, err := os.ReadFile("./maxint.uint64")
	tester.NilErr(t, err)
	res = encode(math.MaxUint64)
	if !bytes.Equal(res, nMax) {
		t.Fatalf("\nWant: %#x\n Got: %#x", nMax, res)
	}
}

func TestDecode(t *testing.T) {

}
