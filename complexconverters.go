package fastconvert

// region Complex64 Converters
// ReadByteArrayToComplex64 reads a 8 byte array to a complex64
func ReadByteArrayToComplex64(data []byte) complex64 {
	var r = ReadByteArrayToFloat32(data[:4])
	var i = ReadByteArrayToFloat32(data[4:8])
	return complex(r, i)
}

// ReadByteArrayToComplex64Array reads a complex64 array from specified byte buffer.
// The number of items are len(data) / 8 or len(out). Which one is the lower.
func ReadByteArrayToComplex64Array(data []byte, out []complex64) {
	var pos = 0
	var itemsOnBuffer = len(data) / 8
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		var r = ReadByteArrayToFloat32(data[:4])
		var i = ReadByteArrayToFloat32(data[4:8])
		out[idx] = complex(r, i)
		pos += 8
	}
}

// ByteArrayToComplex64Array reads a complex64 array from specified byte buffer.
func ByteArrayToComplex64Array(data []byte) []complex64 {
	var out = make([]complex64, len(data) / 8)
	ReadByteArrayToComplex64Array(data, out)
	return out
}
// endregion
// region Complex128 Converters
// ReadByteArrayToComplex128 reads a 16 byte array to a complex128
func ReadByteArrayToComplex128(data []byte) complex128 {
	var r = ReadByteArrayToFloat64(data[:8])
	var i = ReadByteArrayToFloat64(data[8:16])
	return complex(r, i)
}

// ReadByteArrayToComplex128Array reads a complex128 array from specified byte buffer.
// The number of items are len(data) / 16 or len(out). Which one is the lower.
func ReadByteArrayToComplex128Array(data []byte, out []complex128) {
	var pos = 0
	var itemsOnBuffer = len(data) / 16
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		var r = ReadByteArrayToFloat64(data[:8])
		var i = ReadByteArrayToFloat64(data[8:16])
		out[idx] = complex(r, i)
		pos += 8
	}
}

// ByteArrayToComplex128Array reads a complex128 array from specified byte buffer.
func ByteArrayToComplex128Array(data []byte) []complex128 {
	var out = make([]complex128, len(data) / 16)
	ReadByteArrayToComplex128Array(data, out)
	return out
}
// endregion