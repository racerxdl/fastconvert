package fastconvert

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"
)

func TestByteArrayToComplex64(t *testing.T) {
	c := complex(rand.Float32()*2-1, rand.Float32()*2-1)
	bs := new(bytes.Buffer)

	_ = binary.Write(bs, binary.LittleEndian, c)

	c2 := ReadByteArrayToComplex64(bs.Bytes())

	if c2 != c {
		t.Fatalf("Differente number read. Expected %v got %v", c, c2)
	}
}

func TestByteArrayToComplex128(t *testing.T) {
	c := complex(rand.Float64()*2-1, rand.Float64()*2-1)
	bs := new(bytes.Buffer)

	_ = binary.Write(bs, binary.LittleEndian, c)

	c2 := ReadByteArrayToComplex128(bs.Bytes())

	if c2 != c {
		t.Fatalf("Differente number read. Expected %v got %v", c, c2)
	}
}

func TestBinaryReadComplex64Array(t *testing.T) {
	var vecA = make([]complex64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = complex(rand.Float32()*2-1, rand.Float32()*2-1)
	}

	_ = binary.Write(bs, binary.LittleEndian, vecA)

	var vecB = ByteArrayToComplex64Array(bs.Bytes())

	if len(vecB) != len(vecA) {
		t.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			t.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadComplex64ArrayLowOutput(t *testing.T) {
	var vecA = make([]complex64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = complex(rand.Float32()*2-1, rand.Float32()*2-1)
	}

	_ = binary.Write(bs, binary.LittleEndian, vecA)

	vecB := make([]complex64, 8192)

	ReadByteArrayToComplex64Array(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			t.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadComplex128ArrayLowOutput(t *testing.T) {
	var vecA = make([]complex128, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = complex(rand.Float64()*2-1, rand.Float64()*2-1)
	}

	_ = binary.Write(bs, binary.LittleEndian, vecA)

	vecB := make([]complex128, 8192)

	ReadByteArrayToComplex128Array(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			t.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadComplex128Array(b *testing.T) {
	var vecA = make([]complex128, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = complex(rand.Float64()*2-1, rand.Float64()*2-1)
	}

	_ = binary.Write(bs, binary.LittleEndian, vecA)

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
