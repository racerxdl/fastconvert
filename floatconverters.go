package fastconvert


import (
	"encoding/binary"
	"math"
)

// region 32 bits
// ReadByteArrayToFloat32 reads a 4 element byte array and converts to float32
func ReadByteArrayToFloat32(data []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(data))
}

// ReadByteArrayToFloat32Array reads a float32 array from specified byte buffer.
// The number of items are len(data) / 4 or len(out). Which one is the lower.
func ReadByteArrayToFloat32Array(data []byte, out []float32) {
	var pos = 0
	var itemsOnBuffer = len(data) / 8
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = ReadByteArrayToFloat32(data[:4])
		pos += 4
	}
}

// ByteArrayToFloat32Array reads a float32 array from specified byte buffer.
func ByteArrayToFloat32Array(data []byte) []float32 {
	var out = make([]float32, len(data) / 8)
	ReadByteArrayToFloat32Array(data, out)
	return out
}

// endregion
// region 64 bits
// ReadByteArrayToFloat64 reads a 8 element byte array and converts to float64
func ReadByteArrayToFloat64(data []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(data))
}

// ReadByteArrayToFloat64Array reads a float64 array from specified byte buffer.
// The number of items are len(data) / 8 or len(out). Which one is the lower.
func ReadByteArrayToFloat64Array(data []byte, out []float64) {
	var pos = 0
	var itemsOnBuffer = len(data) / 8
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = ReadByteArrayToFloat64(data[:4])
		pos += 4
	}
}

// ByteArrayToFloat64Array reads a float64 array from specified byte buffer.
func ByteArrayToFloat64Array(data []byte) []float64 {
	var out = make([]float64, len(data) / 8)
	ReadByteArrayToFloat64Array(data, out)
	return out
}
// endregion

