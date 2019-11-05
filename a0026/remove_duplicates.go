package a0026

func RemoveDuplicates(nums []int) int {
	length := 0
	if nums != nil {
		size := len(nums)
		if size <= 1 {
			return size
		}
		length = 1
		for i := 0; i < size; i++ {
			if i+1 < size {
				if nums[i] < nums[i+1] {
					nums[length] = nums[i+1]
					length++
				}
			}
		}
	}
	return length
}
