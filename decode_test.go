package lzw

import "testing"

func TestDecode(t *testing.T) {
	encode_vals := []int{97, 97, 98, 256, 97, 100, 256, 98, 115, 261, 262, 265, 266, 259, 97, 115, 257, 256, 100, 261, 258, 260, 259, 275, 267, 115, 258, 271, 277, 283, 256, 281, 269, 257, 285, 270, 272, 271, 293, 270, 283, 98}
	expected_str := "aabaaadaabsdaaabdaaaabdaaaasabaaddabaadaaadabdaaasbasaadasaaaasbaaaaabsaaaasabasasasassaab"
	str_output := Decode(encode_vals)

	if str_output != expected_str {
		t.Errorf(" Expected output of: %s\n but got %s", expected_str, str_output)
	}

}
