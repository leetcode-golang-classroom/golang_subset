package sol

func subsets(nums []int) [][]int {
	result := [][]int{}
	subset := []int{}

	// 透過 dfs
	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			result = append(result, temp)
			return
		}
		// choose i
		subset = append(subset, nums[i])
		dfs(i + 1)
		// not choose i
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}
	dfs(0)
	return result
}
