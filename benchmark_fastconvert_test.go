package fastconvert

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/rand"
	"testing"
)

// region Complex64 Benchmark
func BenchmarkBinaryReadComplex64(b *testing.B) {
	b.StopTimer()
	var vecA = make([]complex64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = complex(rand.Float32()*2-1, rand.Float32()*2-1)
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var b = bytes.NewBuffer(bs.Bytes())
		_ = binary.Read(b, binary.LittleEndian, &vecA)
	}
}
func BenchmarkFastConvertComplex64(b *testing.B) {
	b.StopTimer()
	var vecA = make([]complex64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = complex(rand.Float32()*2-1, rand.Float32()*2-1)
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var b = bs.Bytes()
		ReadByteArrayToComplex64Array(b, vecA)
	}
}
// endregion
// region Complex128 benchmark
func BenchmarkBinaryReadComplex128(b *testing.B) {
	b.StopTimer()
	var vecA = make([]complex128, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = complex(rand.Float64()*2-1, rand.Float64()*2-1)
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var b = bytes.NewBuffer(bs.Bytes())
		_ = binary.Read(b, binary.LittleEndian, &vecA)
	}
}
func BenchmarkFastConvertComplex128(b *testing.B) {
	b.StopTimer()
	var vecA = make([]complex128, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = complex(rand.Float64()*2-1, rand.Float64()*2-1)
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var b = bs.Bytes()
		ReadByteArrayToComplex128Array(b, vecA)
	}
}
// endregion
// region Int16 Benchmark
func BenchmarkBinaryReadInt16(b *testing.B) {
	b.StopTimer()
	var vecA = make([]int16, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int16(rand.Intn(65535) - 32768)
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var b = bytes.NewBuffer(bs.Bytes())
		_ = binary.Read(b, binary.LittleEndian, &vecA)
	}
}
func BenchmarkFastConvertInt16(b *testing.B) {
	b.StopTimer()
	var vecA = make([]int16, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int16(rand.Intn(65535) - 32768)
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var b = bs.Bytes()
		ReadByteArrayToInt16LEArray(b, vecA)
	}
}
// endregion
// region Int32 Benchmark
func BenchmarkBinaryReadInt32(b *testing.B) {
	b.StopTimer()
	var vecA = make([]int32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int32(rand.Intn(math.MaxInt32) - math.MaxInt32 / 2)
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var b = bytes.NewBuffer(bs.Bytes())
		_ = binary.Read(b, binary.LittleEndian, &vecA)
	}
}
func BenchmarkFastConvertInt32(b *testing.B) {
	b.StopTimer()
	var vecA = make([]int32, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int32(rand.Uint32())
	}

	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var b = bs.Bytes()
		ReadByteArrayToInt32LEArray(b, vecA)
	}
}
// endregion
// region Int64 Benchmark
func BenchmarkBinaryReadInt64(b *testing.B) {
	b.StopTimer()
	var vecA = make([]int64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int64(rand.Uint64())
	}


	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var b = bytes.NewBuffer(bs.Bytes())
		_ = binary.Read(b, binary.LittleEndian, &vecA)
	}
}
func BenchmarkFastConvertInt64(b *testing.B) {
	b.StopTimer()
	var vecA = make([]int64, 16384)
	var bs = new(bytes.Buffer)

	for i := 0; i < len(vecA); i++ {
		vecA[i] = int64(rand.Uint64())
	}


	err := binary.Write(bs, binary.LittleEndian, vecA)

	if err != nil {
		panic(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var b = bs.Bytes()
		ReadByteArrayToInt64LEArray(b, vecA)
	}
}
// endregion
