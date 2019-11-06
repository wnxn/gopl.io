package odd

import (
	"fmt"
	"testing"
)

func BenchmarkPrintOdd(b *testing.B) {
	res := PrintOdd(b.N)
	fmt.Println(res)
}
