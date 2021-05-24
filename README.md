## LZW-GO
This is a simple implementation of the LZW compression algorithm in Go

[![Build Status](https://travis-ci.com/j0nimost/lzw.svg?branch=main)](https://travis-ci.com/github/j0nimost/lzw)

### Implementation


Below is an example of the implementation

```
import (
	"fmt"

	"github.com/j0nimost/lzw"
)

func main() {

	// define a data set
	input := "helloworldworldhellow"
	fmt.Println("Your Input: ", input)

	encoded_str := encode.Encode(input)
	decoded_str := decode.Decode(encoded_str)

	fmt.Println(decoded_str)
}
```

### Author
John Nyingi
