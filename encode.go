package lzw

// Encode :(s string)The method encodes string inputs and returns an array int
func Encode(input string) []int {
	ascii := make(map[string]int)
	output := []int{}

	// generate a map of ascii
	for i := 0; i <= 255; i++ {
		char := rune(i)
		str := string(char)
		ascii[str] = i
	}

	// encoding lzw
	compressed := ascii

	r := []rune(input)
	i := 0
	b := string(r[0])
	count := 256

	for i < len(r) {
		if i != len(r)-1 {
			c := string(r[i+1])

			_, ok := compressed[b+c]

			if ok {
				b = b + c
			} else {
				compressed[b+c] = count
				output = append(output, compressed[b])
				b = c
				count++
			}
		}
		i++
	}
	output = append(output, compressed[b])
	return output
}
