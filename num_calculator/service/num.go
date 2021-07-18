package service

// 冒泡排序  从小到大
func bubbleSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if nums[j] < nums[j-1] {
				// 交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

// 是否是质数
func IsPrime(num int) bool {
	for k := 2; k < num; k++ {
		if num%k == 0 {
			return false
		}
	}
	return true
}

// 是否是偶数
func IsEven(num int) bool {
	return num%2 == 0
}
