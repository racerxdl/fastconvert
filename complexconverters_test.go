package fastconvert

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"
)

func TestBinaryReadComplex64(b *testing.T) {
	var vecA = make([]complex64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = complex(rand.Float32()*2-1, rand.Float32()*2-1)
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToComplex64Array(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadComplex128(b *testing.T) {
	var vecA = make([]complex128, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = complex(rand.Float64()*2-1, rand.Float64()*2-1)
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToComplex128Array(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}
