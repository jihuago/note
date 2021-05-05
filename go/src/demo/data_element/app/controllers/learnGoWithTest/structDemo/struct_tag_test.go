package structDemo

import (
	"fmt"
	"testing"
)

func TestNewFile2(t *testing.T) {
	t.Run("test struct tag", func(t *testing.T) {
		tt := TagType{"jack", 1}
		for i := 0; i < 2; i++ {
			refTag(tt, i)
		}
	})

	t.Run("test struct name confict", func(t *testing.T) {
		c := &C{A{1}, B{2, 3}, 4}

		fmt.Println(c.a)

	})
}
