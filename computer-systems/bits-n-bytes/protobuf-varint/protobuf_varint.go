package protobufvarint

func encode(n uint64) []byte {
	res := []byte{}

	for n > 0 {
		// TODO bitmask for possible speed.
		part := n % 128
		// TODO add msb (correctly)
		n >>= 7
		if n > 0 {
			part |= 0b1000_0000
		}
		res = append(res, byte(part))
	}

	return res
}

func decode(bs []byte) uint64 {
	return 0
}
