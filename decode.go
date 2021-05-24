package lzw

// Decode :(arr []int)The method decodes int array and returns a string
func Decode(v []int) string {
	ascii := make(map[int]string)
	outputstr := ""

	// generate a map of ascii
	for i := 0; i <= 255; i++ {
		char := rune(i)

		strc := string(char)
		ascii[i] = strc
	}

	// decode
	decascii := ascii
	n := 0
	old := v[0]
	str := string(decascii[old])
	nchar := ""

	nchar += string([]rune(str)[0])
	asciicount := 256
	outputstr += str

	for b := 0; b < len(v)-1; b++ {
		n = v[b+1]

		_, ok := decascii[n]
		if !ok {
			str = decascii[old]
			str = str + nchar

		} else {
			str = decascii[n]

		}
		outputstr += str
		nchar = ""
		nchar += string(str[0])
		decascii[asciicount] = decascii[old] + nchar
		asciicount++
		old = n
	}

	return outputstr

}
