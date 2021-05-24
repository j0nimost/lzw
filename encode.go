package lzw

func Encode(input string) []int {
	ascii_enc := make(map[string]int)
	output := []int{}

	// generate a map of ascii
	for i := 0; i <= 255; i++ {
		char := rune(i)
		str_c := string(char)
		ascii_enc[str_c] = i
	}

	// encoding lzw
	compressed := ascii_enc

	input_r := []rune(input)
	i := 0
	b := string(input_r[0])
	ascii := 256

	for i < len(input_r) {
		if i != len(input_r)-1 {
			c := string(input_r[i+1])

			_, ok := compressed[b+c]

			if ok {
				b = b + c
			} else {
				compressed[b+c] = ascii
				output = append(output, compressed[b])
				b = c
				ascii++
			}
		}
		i++
	}
	output = append(output, compressed[b])
	return output
}
