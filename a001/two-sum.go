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

/**
通过 map结构，借助kv特性，将另一半提前算好，如果后面遍历时，出现就直接取出
*/
func twoSum1(arr []int, target int) []int {
	var result []int
	var m1 map[int]int
	m1 = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		other := target - arr[i]
		if v, ok := m1[arr[i]]; ok {
			result = append(result, v)
			result = append(result, i)
			return result
		} else {
			m1[other] = i
		}
	}
	return nil
}
