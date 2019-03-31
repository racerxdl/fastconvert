package fastconvert

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"
)

// region 16 bit
func TestBinaryReadInt16LE(b *testing.T) {
	var vecA = make([]int16, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int16(rand.Intn(65535))
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToInt16LEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadInt16LELowOutput(b *testing.T) {
	var vecA = make([]int16, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int16(rand.Intn(65535))
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]int16, 8192)

	ReadByteArrayToInt16LEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadInt16BE(b *testing.T) {
	var vecA = make([]int16, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int16(rand.Intn(65535))
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToInt16BEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadInt16BELowOutput(b *testing.T) {
	var vecA = make([]int16, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int16(rand.Intn(65535))
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]int16, 8192)

	ReadByteArrayToInt16BEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadUInt16LE(b *testing.T) {
	var vecA = make([]uint16, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint16(rand.Intn(65535))
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToUInt16LEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadUInt16LELowOutput(b *testing.T) {
	var vecA = make([]uint16, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint16(rand.Intn(65535))
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]uint16, 8192)

	ReadByteArrayToUInt16LEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadUInt16BE(b *testing.T) {
	var vecA = make([]uint16, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint16(rand.Intn(65535))
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToUInt16BEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadUInt16BELowOutput(b *testing.T) {
	var vecA = make([]uint16, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint16(rand.Intn(65535))
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]uint16, 8192)

	ReadByteArrayToUInt16BEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

// endregion
// region 32 bit
func TestBinaryReadInt32LE(b *testing.T) {
	var vecA = make([]int32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int32(rand.Uint32())
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToInt32LEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadInt32LELowOutput(b *testing.T) {
	var vecA = make([]int32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int32(rand.Uint32())
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]int32, 8182)

	ReadByteArrayToInt32LEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadInt32BE(b *testing.T) {
	var vecA = make([]int32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int32(rand.Uint32())
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToInt32BEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadInt32BELowOutput(b *testing.T) {
	var vecA = make([]int32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int32(rand.Uint32())
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]int32, 8182)

	ReadByteArrayToInt32BEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadUInt32LE(b *testing.T) {
	var vecA = make([]uint32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint32(rand.Uint32())
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToUInt32LEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadUInt32LELowOutput(b *testing.T) {
	var vecA = make([]uint32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint32(rand.Uint32())
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]uint32, 8182)

	ReadByteArrayToUInt32LEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadUInt32BE(b *testing.T) {
	var vecA = make([]uint32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint32(rand.Uint32())
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToUInt32BEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadUInt32BELowOutput(b *testing.T) {
	var vecA = make([]uint32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint32(rand.Uint32())
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]uint32, 8182)

	ReadByteArrayToUInt32BEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

// endregion
// region 64 bit
func TestBinaryReadInt64LE(b *testing.T) {
	var vecA = make([]int64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int64(rand.Uint64())
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToInt64LEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}
func TestBinaryReadInt64LELowOutput(b *testing.T) {
	var vecA = make([]int64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int64(rand.Uint64())
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]int64, 8192)

	ReadByteArrayToInt64LEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadInt64BE(b *testing.T) {
	var vecA = make([]int64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int64(rand.Uint64())
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToInt64BEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}
func TestBinaryReadInt64BELowOutput(b *testing.T) {
	var vecA = make([]int64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int64(rand.Uint64())
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]int64, 8192)

	ReadByteArrayToInt64BEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}
func TestBinaryReadUInt64LE(b *testing.T) {
	var vecA = make([]uint64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint64(rand.Uint64())
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToUInt64LEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}
func TestBinaryReadUInt64LELowOutput(b *testing.T) {
	var vecA = make([]uint64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint64(rand.Uint64())
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]uint64, 8192)

	ReadByteArrayToUInt64LEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

func TestBinaryReadUInt64BE(b *testing.T) {
	var vecA = make([]uint64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint64(rand.Uint64())
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	var vecB = ByteArrayToUInt64BEArray(bs.Bytes())

	if len(vecB) != len(vecA) {
		b.Fatalf("Different array size. Expected %d got %d", len(vecA), len(vecB))
	}

	for i := range vecA {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}
func TestBinaryReadUInt64BELowOutput(b *testing.T) {
	var vecA = make([]uint64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = uint64(rand.Uint64())
	}

	err := binary.Write(bs, binary.BigEndian, vecA)

	if err != nil {
		panic(err)
	}

	vecB := make([]uint64, 8192)

	ReadByteArrayToUInt64BEArray(bs.Bytes(), vecB)

	for i := range vecB {
		if vecA[i] != vecB[i] {
			b.Fatalf("Different array element Expected %v got %v", vecA[i], vecB[i])
		}
	}
}

// endregion
