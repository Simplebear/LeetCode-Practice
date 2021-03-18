package leetcode

func quickSort(nums []int) {
	_quickSort(nums, 0, len(nums)-1)
}

func _quickSort(nums []int, begin int, end int) {
	if begin < end {
		//找到中位数后递归调用
		p := partition(nums, begin, end)
		_quickSort(nums, begin, p-1)
		_quickSort(nums, p+1, end)
	}
}

func partition(nums []int, begin int, end int) int {
	for begin < end {
		//end与begin比较 如果大于begin则end--
		for begin < end && nums[begin] <= nums[end] {
			end--
		}
		// 当end小于begin时交换end和begin位置，begin++
		if begin < end {
			nums[begin], nums[end] = nums[end], nums[begin]
			begin++
		}
		//参考值位置移到end，所以从begin开始向后滑动，如果begin小于end则begin++
		for begin < end && nums[begin] <= nums[end] {
			begin++
		}
		//当begin大于end，交换位置 end--
		if begin < end {
			nums[begin], nums[end] = nums[end], nums[begin]
			end--
		}
		// 交换完位置后如果begin小于end 继续循环
	}
	//循环完毕后 begin变成中间位置
	return begin
}
