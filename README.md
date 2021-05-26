## LZW
This is a simple implementation of the LZW compression algorithm in Go

[![Build Status](https://travis-ci.com/j0nimost/lzw.svg?branch=main)](https://travis-ci.com/github/j0nimost/lzw)

### Implementation
Get the Package

`go get github.com/j0nimost/lzw`

Below is an example of the implementation for

#### Encoding

```
import (
	"fmt"

	"github.com/j0nimost/lzw"
)

func main() {

	// define a data set
	input := "helloworldworldhellow"
	fmt.Println("Your Input: ", input)

	encoded_str := lzw.Encode(input)

	fmt.Printf("Uncompressed: %d\n Compressed: %d\n", len(input), len(encoded_str))
}
```

#### Decoding
```
import (
	"fmt"

	"github.com/j0nimost/lzw"
)

func main() {

	// get a compressed array
	encoded := []int{97, 97, 98, 256, 97, 100, 256, 98, 115, 261, 262, 265, 266, 259, 97, 115, 257, 256, 100, 261, 258, 260, 259, 275, 267, 115, 258, 271, 277, 283, 256, 281, 269, 257, 285, 270, 272, 271, 293, 270, 283, 98}

	decoded_str := lzw.Decode(encoded)

	fmt.Println(decoded_str)
}
```

### Author
John Nyingi
