package fastconvert

import "encoding/binary"

// region 16 bit
// region int16 Converters
// ReadByteArrayToInt16LEArray reads a int16 array from specified byte buffer in Little Endian format.
// The number of items are len(data) / 2 or len(out). Which one is the lower.
func ReadByteArrayToInt16LEArray(data []byte, out []int16) {
	var pos = 0
	var itemsOnBuffer = len(data) / 2
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = int16(binary.LittleEndian.Uint16(data[pos:pos+2]))
		pos += 2
	}
}

func ByteArrayToInt16LEArray(data []byte) []int16 {
	var out = make([]int16, len(data) / 2)
	ReadByteArrayToInt16LEArray(data, out)
	return out
}

// ReadByteArrayToInt16BEArray reads a int16 array from specified byte buffer in Big Endian format.
// The number of items are len(data) / 2 or len(out). Which one is the lower.
func ReadByteArrayToInt16BEArray(data []byte, out []int16) {
	var pos = 0
	var itemsOnBuffer = len(data) / 2
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = int16(binary.BigEndian.Uint16(data[pos:pos+2]))
		pos += 2
	}
}

func ByteArrayToInt16BEArray(data []byte) []int16 {
	var out = make([]int16, len(data) / 2)
	ReadByteArrayToInt16BEArray(data, out)
	return out
}
// endregion
// region uint16 Converters
// ReadByteArrayToInt16LEArray reads a int16 array from specified byte buffer in Little Endian format.
// The number of items are len(data) / 2 or len(out). Which one is the lower.
func ReadByteArrayToUInt16LEArray(data []byte, out []uint16) {
	var pos = 0
	var itemsOnBuffer = len(data) / 2
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = binary.LittleEndian.Uint16(data[pos:pos+2])
		pos += 2
	}
}

func ByteArrayToUInt16LEArray(data []byte) []uint16 {
	var out = make([]uint16, len(data) / 2)
	ReadByteArrayToUInt16LEArray(data, out)
	return out
}

// ReadByteArrayToUInt16BEArray reads a uint16 array from specified byte buffer in Big Endian format.
// The number of items are len(data) / 2 or len(out). Which one is the lower.
func ReadByteArrayToUInt16BEArray(data []byte, out []uint16) {
	var pos = 0
	var itemsOnBuffer = len(data) / 2
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = binary.BigEndian.Uint16(data[pos:pos+2])
		pos += 2
	}
}

// ByteArrayToUInt16BEArray reads a uint16 array from specified byte buffer in Big Endian format.
func ByteArrayToUInt16BEArray(data []byte) []uint16 {
	var out = make([]uint16, len(data) / 2)
	ReadByteArrayToUInt16BEArray(data, out)
	return out
}
// endregion
// endregion
// region 32 bit
// region int32 Converters
// ReadByteArrayToInt32LEArray reads a int32 array from specified byte buffer in Little Endian format.
// The number of items are len(data) / 4 or len(out). Which one is the lower.
func ReadByteArrayToInt32LEArray(data []byte, out []int32) {
	var pos = 0
	var itemsOnBuffer = len(data) / 4
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = int32(binary.LittleEndian.Uint32(data[pos:pos+4]))
		pos += 4
	}
}

func ByteArrayToInt32LEArray(data []byte) []int32 {
	var out = make([]int32, len(data) / 4)
	ReadByteArrayToInt32LEArray(data, out)
	return out
}

// ReadByteArrayToInt32BEArray reads a int32 array from specified byte buffer in Big Endian format.
// The number of items are len(data) / 2 or len(out). Which one is the lower.
func ReadByteArrayToInt32BEArray(data []byte, out []int32) {
	var pos = 0
	var itemsOnBuffer = len(data) / 4
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = int32(binary.BigEndian.Uint32(data[pos:pos+4]))
		pos += 4
	}
}

func ByteArrayToInt32BEArray(data []byte) []int32 {
	var out = make([]int32, len(data) / 4)
	ReadByteArrayToInt32BEArray(data, out)
	return out
}
// endregion
// region uint32 Converters
// ReadByteArrayToInt32LEArray reads a int32 array from specified byte buffer in Little Endian format.
// The number of items are len(data) / 4 or len(out). Which one is the lower.
func ReadByteArrayToUInt32LEArray(data []byte, out []uint32) {
	var pos = 0
	var itemsOnBuffer = len(data) / 4
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = binary.LittleEndian.Uint32(data[pos:pos+4])
		pos += 4
	}
}

func ByteArrayToUInt32LEArray(data []byte) []uint32 {
	var out = make([]uint32, len(data) / 4)
	ReadByteArrayToUInt32LEArray(data, out)
	return out
}

// ReadByteArrayToUInt32BEArray reads a uint32 array from specified byte buffer in Big Endian format.
// The number of items are len(data) / 2 or len(out). Which one is the lower.
func ReadByteArrayToUInt32BEArray(data []byte, out []uint32) {
	var pos = 0
	var itemsOnBuffer = len(data) / 4
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = binary.BigEndian.Uint32(data[pos:pos+4])
		pos += 4
	}
}

// ByteArrayToUInt32BEArray reads a uint32 array from specified byte buffer in Big Endian format.
func ByteArrayToUInt32BEArray(data []byte) []uint32 {
	var out = make([]uint32, len(data) / 4)
	ReadByteArrayToUInt32BEArray(data, out)
	return out
}
// endregion
// endregion
// region 64 bit
// region int64 Converters
// ReadByteArrayToInt64LEArray reads a int64 array from specified byte buffer in Little Endian format.
// The number of items are len(data) / 8 or len(out). Which one is the lower.
func ReadByteArrayToInt64LEArray(data []byte, out []int64) {
	var pos = 0
	var itemsOnBuffer = len(data) / 8
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = int64(binary.LittleEndian.Uint64(data[pos:pos+8]))
		pos += 8
	}
}

func ByteArrayToInt64LEArray(data []byte) []int64 {
	var out = make([]int64, len(data) / 8)
	ReadByteArrayToInt64LEArray(data, out)
	return out
}

// ReadByteArrayToInt64BEArray reads a int64 array from specified byte buffer in Big Endian format.
// The number of items are len(data) / 8 or len(out). Which one is the lower.
func ReadByteArrayToInt64BEArray(data []byte, out []int64) {
	var pos = 0
	var itemsOnBuffer = len(data) / 8
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = int64(binary.BigEndian.Uint64(data[pos:pos+8]))
		pos += 8
	}
}

func ByteArrayToInt64BEArray(data []byte) []int64 {
	var out = make([]int64, len(data) / 8)
	ReadByteArrayToInt64BEArray(data, out)
	return out
}
// endregion
// region uint64 Converters
// ReadByteArrayToInt64LEArray reads a int64 array from specified byte buffer in Little Endian format.
// The number of items are len(data) / 8 or len(out). Which one is the lower.
func ReadByteArrayToUInt64LEArray(data []byte, out []uint64) {
	var pos = 0
	var itemsOnBuffer = len(data) / 8
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = binary.LittleEndian.Uint64(data[pos:pos+8])
		pos += 8
	}
}

func ByteArrayToUInt64LEArray(data []byte) []uint64 {
	var out = make([]uint64, len(data) / 8)
	ReadByteArrayToUInt64LEArray(data, out)
	return out
}

// ReadByteArrayToUInt64BEArray reads a uint64 array from specified byte buffer in Big Endian format.
// The number of items are len(data) / 8 or len(out). Which one is the lower.
func ReadByteArrayToUInt64BEArray(data []byte, out []uint64) {
	var pos = 0
	var itemsOnBuffer = len(data) / 8
	var itemsToRead = itemsOnBuffer
	if len(out) < itemsToRead {
		itemsToRead = len(out)
	}

	for idx := 0; idx < itemsToRead; idx++ {
		out[idx] = binary.BigEndian.Uint64(data[pos:pos+8])
		pos += 8
	}
}

// ByteArrayToUInt64BEArray reads a uint64 array from specified byte buffer in Big Endian format.
func ByteArrayToUInt64BEArray(data []byte) []uint64 {
	var out = make([]uint64, len(data) / 8)
	ReadByteArrayToUInt64BEArray(data, out)
	return out
}
// endregion
// endregion
