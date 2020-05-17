package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	cases := [][]int{
		[]int{10, 2, 4, 6, 5, 2, 1},
		[]int{10, 1},
		[]int{},
	}
	for _, c := range cases {
		length := len(c)
		MergeSort(c, 0, len(c))
		fmt.Println(c)
		assert.Len(t, c, length)
		for i := range c {
			if i == 0 {
				continue
			}
			assert.LessOrEqual(t, c[i-1], c[i])
		}
	}
}
