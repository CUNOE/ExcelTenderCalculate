package core

func IsContain(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

func GenerateNums(n int) []int {
	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	return nums
}

func AvgPrice(nums []float64) float64 {
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}

func CalculateRandomPrices(avgprice float64, randomValue []float64) []float64 {
	var RandomPrices []float64
	for _, v := range randomValue {
		RandomPrices = append(RandomPrices, avgprice*(1+(v*0.01)))
	}
	return RandomPrices
}

func Zhongbiao(randomPrice float64, companys []Company) Company {
	var de []float64

	for _, v := range companys {
		de = append(de, v.Price-randomPrice)
	}

	if IsPositiveNums(de) {
		d := get_zuijin(0, de)

		price := randomPrice + d

		for _, v := range companys {
			if v.Price == price {
				return v
			}
		}
	} else {
		d := zuijin_Minus(de)
		price := randomPrice + d
		for _, v := range companys {
			if v.Price == price {
				return v
			}
		}
	}

	return Company{}
}

func IsPositiveNums(nums []float64) bool {
	for _, n := range nums {
		if n < 0 {
			return false
		}
	}
	return true
}

func zuijin_Minus(nums []float64) float64 {
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

func get_zuijin(this float64, arr []float64) float64 {
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
