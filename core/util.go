package core

import "sort"

// DeleteSlice 从切片中删除 0 元素
func DeleteSlice(a []int) []int {
	ret := make([]int, 0, len(a))
	//var ret []int
	for _, val := range a {
		if val != 0 {
			ret = append(ret, val)
		}
	}
	return ret
}

// RemoveDuplicatesInPlace 删除切片中的重复元素
func RemoveDuplicatesInPlace(userIDs []int) []int {
	// 如果有0或1个元素，则返回切片本身。
	if len(userIDs) < 2 {
		return userIDs
	}

	//  使切片升序排序
	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })

	uniqPointer := 0

	for i := 1; i < len(userIDs); i++ {
		// 比较当前元素和唯一指针指向的元素
		//  如果它们不相同，则将项写入唯一指针的右侧。
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]
}

// IsContain 查询切片中是否存在某个元素
func IsContain(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

// GenerateNums 根据n生成有序数组
func GenerateNums(n int) []int {
	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	return nums
}

// AvgPrice 计算平均值
func AvgPrice(nums []float64) float64 {
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}

// CalculateRandomPrices 计算随机相对应价格
func CalculateRandomPrices(avgPrice float64, randomValue []float64) []float64 {
	var RandomPrices []float64
	for _, v := range randomValue {
		RandomPrices = append(RandomPrices, avgPrice*(1+(v*0.01)))
	}
	return RandomPrices
}

// ReturnCompanyWhoWinTheBidding 计算在某个随机相对应价格下中标的公司
func ReturnCompanyWhoWinTheBidding(randomPrice float64, companies []Company) Company {
	var de []float64

	for _, v := range companies {
		de = append(de, v.Price-randomPrice)
	}

	if IsPositiveNums(de) {
		d := numClosestToZeroPositiveNums(0, de)

		price := randomPrice + d

		for _, v := range companies {
			if v.Price == price {
				return v
			}
		}
	} else {
		d := numClosestToZeroWithMinus(de)
		price := randomPrice + d
		for _, v := range companies {
			if v.Price == price {
				return v
			}
		}
	}

	return Company{}
}

// IsPositiveNums 检查数组中是否全为正数
func IsPositiveNums(nums []float64) bool {
	for _, n := range nums {
		if n < 0 {
			return false
		}
	}
	return true
}

// numClosestToZeroWithMinus 返回数组中最接近0的负数
func numClosestToZeroWithMinus(nums []float64) float64 {
	max := 0.0
	for _, v := range nums {
		if v < 0.0 {
			if max == 0.0 {
				max = v
			}
			if v > max {
				max = v
			}
		}
	}
	return max
}

// numClosestToZero_PositiveNums 返回数组中最接近0的正数
func numClosestToZeroPositiveNums(this float64, arr []float64) float64 {
	min := 0.0
	if this == arr[0] {
		return arr[0]
	} else if this > arr[0] {
		min = this - arr[0]
	} else if this < arr[0] {
		min = arr[0] - this
	}

	for _, v := range arr {
		if v == this {
			return v
		} else if v > this {
			if min > v-this {
				min = v - this
			}
		} else if v < this {
			if min > this-v {
				min = this - v
			}
		}
	}

	for _, v := range arr {
		if this+min == v {
			return v
		} else if this-min == v {
			return v
		}
	}
	return min
}
