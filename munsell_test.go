package munsell

import (
	"testing"
	"fmt"
)

func TestConvertBetween(t *testing.T) {
	fmt.Println(ConvertBetween("#ffffff"))
	fmt.Println(ConvertBetween("2.5R 1:2"))
}
