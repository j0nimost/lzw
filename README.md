## LZW-GO
This is a simple implementation of the LZW compression algorithm in Go

### Implementation

Below is an example of the implimentation

```
import (
	"fmt"

	"github.com/j0nimost/lzw_go/decode"
	"github.com/j0nimost/lzw_go/encode"
)

func main() {

	// define a data set
	input := "helloworldworldhellow"
	fmt.Println("Your Input: ", input)
	
	enc_values := encode.LZW_Encode(input)
	decoded_str := decode.LZW_Decode(enc_values)

	fmt.Println(decoded_str)
}
```

### Author
John Nyingi
