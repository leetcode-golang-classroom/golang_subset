# golang_subset

## 

Given an integer array `nums` of **unique** elements, return *all possible subsets (the power set)*.

The solution set **must not** contain duplicate subsets. Return the solution in **any order**.

## Examples

**Example 1:**

```
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

```

**Example 2:**

```
Input: nums = [0]
Output: [[],[0]]

```

**Constraints:**

- `1 <= nums.length <= 10`
- `10 <= nums[i] <= 10`
- All the numbers of `nums` are **unique**.

## 解析

題目給定一組整數陣列 nums ，

求寫出一個演算法來推算出所有可能的陣列子集合

對於每個元素都為可選或可不選

假設有 n 個元素，就有 $2^n$ 個子集合

如下圖：

![](https://i.imgur.com/puuBL8Z.png)

這樣的話就可以透過 dfs 跟 prune來做到找出所有可能性了

## 程式碼

```go
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
```
## 困難點

1. 需要理解窮舉法的規則
2. 透過一些規則來窮舉
3. 需要理解 golang slice 只是一個 reference 到 array 所以必須 copy 來儲存當下值

## Solve Point

- [x]  Understand what problem to solve
- [x]  Analysis Complexity