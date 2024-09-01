package revisitingarraysandsliceswithgenerics

func Sum(nums []int) int {
	return Reduce(nums, func(acc int, v int) int {
		return acc + v
	}, 0)
}

func SumAllTails(numsToSum ...[]int) []int {
	return Reduce(numsToSum, func(acc []int, v []int) []int {
		if len(v) == 0 {
			acc = append(acc, 0)
		} else {
			acc = append(acc, Sum(v[1:]))
		}
		return acc
	}, make([]int, 0, len(numsToSum)))
}
