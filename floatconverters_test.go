package fastconvert

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"
)

func TestBinaryReadFloat32(b *testing.T) {
	var vecA = make([]float32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = rand.Float32()*2 - 1
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToFloat32Array(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadFloat64(b *testing.T) {
	var vecA = make([]float64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = rand.Float64()*2 - 1
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToFloat64Array(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadFloat64LowOutput(b *testing.T) {
	var vecA = make([]float64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = rand.Float64()*2 - 1
	}

	_ = binary.Write(bs, binary.LittleEndian, vecA)

	vecB := make([]float64, 8192)

	ReadByteArrayToFloat64Array(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadFloat32LowOutput(b *testing.T) {
	var vecA = make([]float32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = rand.Float32()*2 - 1
	}

	_ = binary.Write(bs, binary.LittleEndian, vecA)

	vecB := make([]float32, 8192)

	ReadByteArrayToFloat32Array(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}
