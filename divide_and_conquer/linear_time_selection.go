package divide_and_conquer

import "math"

const NumOfElementsInOneGroup = 5

func linearTimeSelection(nums []int, i, j, k int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	l, r := partition(nums, i, j, getMedian(nums[i:j+1]))

	if k >= l && k <= r {
		return nums[l]
	} else if k < l {
		return linearTimeSelection(nums, i, l-1, k)
	} else {
		return linearTimeSelection(nums, r+1, j, k)
	}
}

func getMedian(nums []int) int {
	if len(nums) == 0 {
		panic("the length of nums should be greater than 0")
	}
	medians := make([]int, int(math.Ceil(float64(len(nums)) / NumOfElementsInOneGroup)))
	for i := range medians {
		start, end := i*NumOfElementsInOneGroup, (i+1)*NumOfElementsInOneGroup
		if end > len(nums) {
			end = len(nums)
		}
		insertSort(nums[start:end])
		medians[i] = nums[(start+end)>>1]
	}
	if len(medians) == 1 {
		return medians[0]
	}
	return getMedian(medians)
}

func insertSort(nums []int) {
	for i:=1; i<len(nums); i++ {
		for j:=i-1; j>=0; j-- {
			if nums[j+1] >= nums[j] {
				break
			}
			nums[j+1], nums[j] = nums[j], nums[j+1]
		}
	}
}

func partition(nums []int, i, j, pivot int) (int, int) {
	l, r := i, j
	for cnt:=i; cnt <=r; {
		if nums[cnt] > pivot {
			nums[cnt], nums[r] = nums[r], nums[cnt]
			r--
		} else if nums[cnt] < pivot {
			nums[cnt], nums[l] = nums[l], nums[cnt]
			l++
			cnt++
		} else {
			cnt++
		}
	}
	return l, r
}
