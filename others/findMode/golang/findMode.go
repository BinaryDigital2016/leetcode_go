package others

// 1.暴力法：两层循环挨个计算每个数字的频数是否超过一半
// 2.hash法：记录所有数字的频度，找出超过一半的
// 3.排序法：将nums排序，第n/2（n/2+1）个一定是众数

// 4.投票法
// 想法
// 如果我们把众数记为 +1+1+1 ，把其他数记为 −1-1−1 ，将它们全部加起来，显然和大于 0 ，从结果本身我们可以看出众数比其他数多。
// 算法
// 本质上， Boyer-Moore 算法就是找 nums 的一个后缀 sufsufsuf ，其中 suf[0]suf[0]suf[0] 就是后缀中的众数。我们维护一个计数器，如果遇到一个我们目前的候选众数，就将计数器加一，否则减一。只要计数器等于 0 ，我们就将 nums 中之前访问的数字全部 忘记 ，并把下一个数字当做候选的众数。直观上这个算法不是特别明显为何是对的，我们先看下面这个例子（竖线用来划分每次计数器归零的情况）
// [7, 7, 5, 7, 5, 1 | 5, 7 | 5, 5, 7, 7 | 7, 7, 7, 7]
// 首先，下标为 0 的 7 被当做众数的第一个候选。在下标为 5 处，计数器会变回0 。所以下标为 6 的 5 是下一个众数的候选者。由于这个例子中 7 是真正的众数，所以通过忽略掉前面的数字，我们忽略掉了同样多数目的众数和非众数。因此， 7 仍然是剩下数字中的众数。
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	ret := nums[0]
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			ret = nums[i]
		}

		if ret == nums[i] {
			count++
		} else {
			count--
		}
	}
	return ret
}
