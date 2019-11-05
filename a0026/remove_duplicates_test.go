package a0026

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Printf("nums = %v \n", nums)
	length := RemoveDuplicates(nums)
	fmt.Printf("nums = %v length=%v \n", nums, length)
}
