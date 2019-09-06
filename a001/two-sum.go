package a001

func twoSum(arr []int, target int) []int {
	var result []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			} else {
				if (arr[i] + arr[j]) == target {
					result = append(result, i)
					result = append(result, j)
					return result
				}
			}
		}
	}
	return nil
}

func twoSum1(arr []int, target int) []int {
	//var result []int
	for i := 0; i < len(arr); i++ {

	}
	return nil
}
