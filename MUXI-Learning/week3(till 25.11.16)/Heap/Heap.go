
func heapCountIterative(nums []int) [][]int {
    var result [][]int
    LEN := len(nums)
    count := make([]int, LEN) 

    result = append(result, append([]int(nil), nums...))

    i := 0
    for i < LEN {
        if count[i] < i { 
            if (i+1)%2 == 0 { 
                nums[count[i]], nums[i] = nums[i], nums[count[i]]
            } else { 
                nums[0], nums[i] = nums[i], nums[0]
            }

            result = append(result, append([]int(nil), nums...))

            count[i]++
            i = 0
        } else {
            count[i] = 0
            i++
        }
    }

    return result
}

func permute(nums []int) [][]int {
	var result [][]int 
	LEN := len(nums)

	tempNums := make([]int, LEN)
	copy(tempNums, nums)

	var heapRec func(k int)

	heapRec = func(k int) {
		if k == 1 {
			curPerm := make([]int, LEN)
			copy(curPerm, tempNums)
			result = append(result, curPerm)
			return
		}

		for i := 0; i < k; i++ {
			heapRec(k - 1) 
			if k%2 == 0 {
				tempNums[i], tempNums[k-1] = tempNums[k-1], tempNums[i]
			} else {
				tempNums[0], tempNums[k-1] = tempNums[k-1], tempNums[0]
			}
		}
	}

	heapRec(LEN)

	return result
}