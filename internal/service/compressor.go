package service

func Compress(data []byte) []byte {
	// Dummy implementation of compression
	return data
}

func Decompress(data []byte) []byte {
	// Dummy implementation of decompression
	return data
}

func CharFrequency(s string) map[rune]int {
	freq := make(map[rune]int)

	for _, ch := range s {
		freq[ch]++
	}

	return freq
}
