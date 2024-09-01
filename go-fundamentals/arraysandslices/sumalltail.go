package arraysandslices

func SumAllTails(numsToSum ...[]int) []int {
	result := make([]int, 0, len(numsToSum))

	for _, nums := range numsToSum {
		if len(nums) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(nums[1:]))
		}
	}

	return result
}
