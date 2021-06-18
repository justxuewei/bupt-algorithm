package divide_and_conquer

import (
	"algorithm/foundation"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestInsertSort(t *testing.T) {
	case1 := []int{5, 2, 4, 3, 1}
	insertSort(case1)
	t.Log(case1)
}

func TestLinearTimeSelection(t *testing.T) {
	radNums := generateNumbersRandomly(0, 100000, 10000)
	srtNums := deepCopy(radNums)
	sort.Ints(srtNums)
	// case1
	case1 := deepCopy(radNums)
	case1Sorted := deepCopy(srtNums)
	for i := range case1 {
		assert.Equal(t, case1Sorted[i], linearTimeSelection(case1, 0, len(case1)-1, i))
	}
	// case2
	case2 := deepCopy(radNums)
	pendingSelection := generateNumbersRandomly(0, 10000, 200)
	maxLen := 0
	for _, k := range pendingSelection {
		linearTimeSelection(case2, 0, len(case2)-1, k)
		continuousLen := 1
		for j:=1; j<len(case2); j++ {
			if case2[j-1] < case2[j] {
				continuousLen++
			} else {
				maxLen = foundation.Max(maxLen, continuousLen)
				continuousLen = 1
			}
		}
		t.Logf("maximum continuous length == %d", maxLen)
	}
}

func generateNumbersRandomly(l, r, n int) []int {
	nums := make([]int, n)
	rad := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range nums {
		//生成随机数
		nums[i] = rad.Intn(r - l) + l
	}
	return nums
}

func deepCopy(src []int) []int {
	dst := make([]int, len(src))
	for i := range src {
		dst[i] = src[i]
	}
	return dst
}
