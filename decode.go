package lzw

func Decode(v []int) string {
	ascii_deco := make(map[int]string)
	output_str := ""

	// generate a map of ascii
	for i := 0; i <= 255; i++ {
		char := rune(i)

		str_c := string(char)
		ascii_deco[i] = str_c
	}

	// decode
	dec_ascii := ascii_deco
	n := 0
	old := v[0]
	str := string(dec_ascii[old])
	n_char := ""

	n_char += string([]rune(str)[0])
	ascii_count := 256
	output_str += str

	for b := 0; b < len(v)-1; b++ {
		n = v[b+1]

		_, ok := dec_ascii[n]
		if !ok {
			str = dec_ascii[old]
			str = str + n_char

		} else {
			str = dec_ascii[n]

		}
		output_str += str
		n_char = ""
		n_char += string(str[0])
		dec_ascii[ascii_count] = dec_ascii[old] + n_char
		ascii_count++
		old = n
	}

	return output_str

}
