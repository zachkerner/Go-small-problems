package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func longestPalindrome(s string) string {
	l, r := 0, len(s)-1
	return helper(s, l, r)
}

func helper(s string, l int, r int) string {

	if l == r {
		return string(s[l])
	}

	if l > r {
		return ""
	}

	if s[l] == s[r] {
		return string(s[l]) + helper(s, l+1, r-1) + string(s[l])
	}

	return max(helper(s, l+1, r), helper(s, l, r-1))
}

func max(x, y string) string {
	if len(x) >= len(y) {
		return x
	}
	return y
}

func calculate(s string) int {
	stack := []string{}

	numStr := "0123456789"
	current := ""
	for i := 0; i < len(s); i++ {
		//fmt.Println(current, stack)
		char := string(s[i])
		if strings.Contains(numStr, char) {
			current += char
		}

		if char == "+" || char == "-" || char == "/" || char == "*" {
			if len(stack) <= 1 {
				stack = append(stack, current)
				current = ""
				continue
			}

			if stack[len(stack)-1] == "*" || stack[len(stack)-1] == "/" {
				firstOperand, _ := strconv.Atoi(stack[len(stack)-2])
				secondOperand, _ := strconv.Atoi(current)
				result := 0
				if stack[len(stack)-1] == "*" {
					result = firstOperand * secondOperand
				}
				if stack[len(stack)-1] == "/" {
					result = firstOperand / secondOperand
				}

				current = strconv.Itoa(result)
				stack = stack[:len(stack)-2]
			}
			stack = append(stack, current, char)
			current = ""
		}

		if i == len(s)-1 && len(current) > 0 {
			stack = append(stack, current)
		}
	}
	fmt.Println(stack)
	return 0
}

func main() {
	fmt.Println(calculate("3/2"))
}

// func findMin(nums []int) int {
// 	left, right := 0, len(nums)-1

// 	if nums[right] >= nums[left] {
// 		return nums[0]
// 	}

// 	for left <= right {
// 		mid := (left + right) / 2
// 		val := isPeakOrValley(nums, mid)

// 		if val == -1 {
// 			return nums[mid]
// 		}

// 		if val == 1 {
// 			return nums[mid+1]
// 		}

// 		if nums[right] >= nums[mid] {
// 			right = mid - 1
// 			continue
// 		}

// 		if nums[left] <= nums[mid] {
// 			left = mid + 1
// 			continue
// 		}
// 	}

// 	return nums[left]
// }

// func isPeakOrValley(nums []int, idx int) int {
// 	if idx == 0 || idx == len(nums)-1 {
// 		return 0
// 	}

// 	if nums[idx] > nums[idx+1] {
// 		return 1
// 	}

// 	if nums[idx] < nums[idx-1] {
// 		return -1
// 	}
// 	return 0
// }

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 0

	for i < len(intervals)-1 {
		firstLow, firstHigh := intervals[i][0], intervals[i][1]
		secondLow, secondHigh := intervals[i+1][0], intervals[i+1][1]

		if firstHigh >= secondLow {
			mergedArr := []int{firstLow, secondHigh}

			if secondHigh < firstHigh {
				mergedArr[0], mergedArr[1] = firstLow, firstHigh
			}

			intervals[i] = mergedArr
			intervals = append(intervals[0:i+1], intervals[(i+2):len(intervals)]...)

			continue
		}

		i++
	}

	return intervals
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, v := range nums {
		hash[v] = i
	}

	for idx, elem := range nums {

		sum := target - elem
		if hash[sum] > 0 {
			return []int{idx, hash[elem]}
		}
	}

	return []int{0, 0}
}

// func removeDuplicates(nums []int) int {
// 	a, r := 0, 1

// 	for r < len(nums) {
// 		if nums[a] == nums[r] {
// 			nums[r] = -1001
// 			r++
// 			continue
// 		}

// 		if nums[a] < nums[r] {
// 			a = r
// 			r++
// 		}
// 	}

// 	a = 0
// 	r = 0

// 	for r < len(nums) {
// 		if nums[r] != -1001 {
// 			nums[a], nums[r] = nums[r], nums[a]
// 			a++
// 		}

// 		r++
// 	}

// 	return a
// }

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	hashIdx := make(map[int][]int)
	hashDup := make(map[string]bool)

	for idx, v := range nums {
		if len(hashIdx[v]) > 0 {
			hashIdx[v] = append(hashIdx[v], idx)
		}
		if len(hashIdx[v]) == 0 {
			hashIdx[v] = []int{idx}
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := (nums[i] + nums[j]) * -1
			if len(hashIdx[sum]) > 0 {
				for k := 0; k < len(hashIdx[sum]); k++ {
					if hashIdx[sum][k] == i || hashIdx[sum][k] == j {
						continue
					}
					idxArr := []int{nums[i], nums[j], nums[hashIdx[sum][k]]}
					idxStr := makeStr(idxArr)

					if !hashDup[idxStr] {
						result = append(result, idxArr)
						hashDup[idxStr] = true
					}
				}
			}
		}
	}
	return result
}

func makeStr(idxArr []int) string {
	sort.Slice(idxArr, func(i, j int) bool {
		return idxArr[i] < idxArr[j]
	})
	idxStr := ""
	for _, num := range idxArr {
		idxStr += strconv.Itoa(num)
	}
	return idxStr
}

func isPalindrome(s string) bool {
	regStr := strings.Join(regexp.MustCompile("([[:alpha:]]|\\d)*").FindAllString(strings.ToLower(s), -1), "")

	if len(regStr) <= 1 {
		return true
	}

	lo := regStr[0]
	hi := regStr[len(regStr)-1]
	fmt.Println(string(lo), string(hi))

	return (lo == hi) && isPalindrome(regStr[1:len(regStr)-1])
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		fmt.Println(left, right, mid)
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[right] { //we're on a binary line
			if nums[mid] < target && target <= nums[right] {
				return binarySearch(nums, target, mid+1, right)
			}
			right = mid
			continue
		}

		if nums[mid] > nums[right] { //left to mid is a binary line
			if nums[left] < target && target < nums[mid] {
				return binarySearch(nums, target, left, mid-1)
			}
			left = mid + 1
		}
	}
	return -1
}

func binarySearch(nums []int, target int, leftIdx int, rightIdx int) int {
	left, right := leftIdx, rightIdx

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

// func maxSubArray(nums []int) int {
// 	if len(nums) < 1 {
// 		return int(math.Inf(-1))
// 	}
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	left, right := 0, len(nums)-1
// 	mid := (left + right) / 2 //average

// 	leftSum := maxSubArray(nums[:mid+1])
// 	rightSum := maxSubArray(nums[mid+1:])
// 	crossSum := crossSum(nums)

// 	return max(leftSum, rightSum, crossSum)
// }

// func crossSum(nums []int) int {
// 	leftSum := int(math.Inf(-1))
// 	rightSum := int(math.Inf(-1))

// 	l, r := 0, len(nums)-1
// 	m := (r + l) / 2

// 	curSumL := 0
// 	for i := m; i >= 0; i-- {
// 		curSumL += nums[i]
// 		if curSumL > leftSum {
// 			leftSum = curSumL
// 		}
// 	}

// 	curSumR := 0
// 	for j := m + 1; j < len(nums); j++ {
// 		curSumR += nums[j]
// 		if curSumR > rightSum {
// 			rightSum = curSumR
// 		}
// 	}

// 	return max(leftSum, rightSum, leftSum+rightSum)
// }

// func max(l, r, c int) int {
// 	maxEl := int(math.Inf(-1))
// 	arr := []int{l, r, c}
// 	for _, el := range arr {
// 		if el > maxEl {
// 			maxEl = el
// 		}
// 	}
// 	return maxEl
// }

// func main() {
// 	arr := []int{5, 4, -1, 7, 8}
// 	fmt.Println(maxSubArray(arr))
// }

// // func search(nums []int, target int) int {
// //   minIdx := findMin(nums)

// //   toMin := binarySearch(nums[:minIdx], target)
// //   postMin := binarySearch(nums[minIdx:], target)

// //   if postMin != -1 {
// //     return postMin + len(nums[:minIdx])
// //   }

// //   return toMin
// // }

// // func binarySearch(nums []int, target int) int {
// //   left, right := 0, len(nums)-1

// //   for left <= right {
// //     mid := (left + right)/2
// //     if nums[mid] == target {
// //       return mid
// //     }

// //     if nums[mid] < target {
// //       left = mid + 1
// //     }

// //     if nums[mid] > target {
// //       right = mid - 1
// //     }
// //   }
// //   return -1
// // }

// func findMin(nums []int) int {
// 	left, right := 0, len(nums)-1

// 	if nums[right] >= nums[left] {
// 		return nums[0]
// 	}

// 	for left <= right {
// 		mid := (left + right) / 2
// 		val := isPeakOrValley(nums, mid)

// 		if val == -1 {
// 			return mid
// 		}

// 		if val == 1 {
// 			return mid + 1
// 		}

// 		if nums[right] >= nums[mid] {
// 			right = mid - 1
// 			continue
// 		}

// 		if nums[left] <= nums[mid] {
// 			left = mid + 1
// 			continue
// 		}
// 	}

// 	return left
// }

// func isPeakOrValley(nums []int, idx int) int {
// 	if idx == 0 || idx == len(nums)-1 {
// 		return 0
// 	}

// 	if nums[idx] > nums[idx+1] {
// 		return 1
// 	}

// 	if nums[idx] < nums[idx-1] {
// 		return -1
// 	}
// 	return 0
// }

// func minimumTotal(triangle [][]int) int {

// 	return helper(triangle, 0, 0)
// }

// func helper(triangle [][]int, row int, idx int) int {
// 	fmt.Println("row", row, "idx", idx, len(triangle)-1)
// 	if row >= len(triangle) || idx >= len(triangle[row]) {
// 		return 0
// 	}
// 	if row == len(triangle)-1 {
// 		return triangle[row][idx]
// 	}

// 	secondNum := helper(triangle, row+1, idx)
// 	thirdNum := helper(triangle, row+1, idx+1)

// 	return min(triangle[row][idx]+secondNum, triangle[row][idx]+thirdNum)
// }

// func min(x, y int) int {
// 	if x <= y {
// 		return x
// 	} else {
// 		return y
// 	}
// 	return x
// }

// func coinChange(coins []int, amount int) int {
// 	count := 0
// 	current := amount
// 	sort.Ints(coins)
// 	return helper(coins, amount, current, count)
// }

// func helper(coins []int, amount int, current int, count int) int {
// 	if len(coins) <= 0 {

// 		return -1
// 	}
// 	//fmt.Println(current, coins)

// 	curCoin := coins[len(coins)-1]

// 	if current%curCoin == 0 {
// 		return current / curCoin
// 	}

// 	if current-curCoin < 0 {
// 		return helper(coins[:len(coins)-1], amount, current, count)
// 	}

// 	maxVal := (current / curCoin) * curCoin

// 	for i := maxVal; i >= curCoin; i -= curCoin {
// 		if len(coins) < 0 {
// 			break
// 		}

// 		result := helper(coins[:len(coins)-1], amount, current-i, count)

// 		if result == -1 {
// 			continue
// 		}

// 		if result != -1 {

// 			return result + (current / curCoin)
// 		}

// 	}

// 	newCoins := coins[:len(coins)-1]
// 	return helper(newCoins, amount, amount, 0)
// }
