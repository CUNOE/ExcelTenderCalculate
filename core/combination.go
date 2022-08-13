package core

import (
	"log"
)

// Combination 组合算法
func Combination(nums []int, m int) [][]int {

	n := len(nums)
	indexes := combinationResult(n, m)
	result := findNumbsByIndexes(nums, indexes)
	//结果是否正确
	rightCount := mathCombination(n, m)
	if rightCount == len(result) {
		//fmt.Println("结果正确")
		log.Printf("组合结果正确")
	} else {
		//fmt.Println("结果错误，正确结果是：", rightCount)
		log.Fatalf("组合结果错误，正确结果是：%d", rightCount)
	}

	return result
}

//组合算法(从nums中取出m个数)
func combinationResult(n int, m int) [][]int {
	if m < 1 || m > n {
		//fmt.Println("Illegal argument. Param m must between 1 and len(nums).")
		log.Fatalf("Illegal argument. Param m must between 1 and len(nums).")
		return [][]int{}
	}
	//保存最终结果的数组，总数直接通过数学公式计算
	result := make([][]int, 0, mathCombination(n, m))
	//保存每一个组合的索引的数组，1表示选中，0表示未选中
	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		if i < m {
			indexes[i] = 1
		} else {
			indexes[i] = 0
		}
	}
	//第一个结果
	result = addTo(result, indexes)
	for {
		find := false
		//每次循环将第一次出现的 1 0 改为 0 1，同时将左侧的1移动到最左侧
		for i := 0; i < n-1; i++ {
			if indexes[i] == 1 && indexes[i+1] == 0 {
				find = true
				indexes[i], indexes[i+1] = 0, 1
				if i > 1 {
					moveOneToLeft(indexes[:i])
				}
				result = addTo(result, indexes)
				break
			}
		}
		//本次循环没有找到 1 0 ，说明已经取到了最后一种情况
		if !find {
			break
		}
	}
	return result
}

//将ele复制后添加到arr中，返回新的数组
func addTo(arr [][]int, ele []int) [][]int {
	newEle := make([]int, len(ele))
	copy(newEle, ele)
	arr = append(arr, newEle)
	return arr
}
func moveOneToLeft(leftNums []int) {
	//计算有几个1
	sum := 0
	for i := 0; i < len(leftNums); i++ {
		if leftNums[i] == 1 {
			sum++
		}
	}
	//将前sum个改为1，之后的改为0
	for i := 0; i < len(leftNums); i++ {
		if i < sum {
			leftNums[i] = 1
		} else {
			leftNums[i] = 0
		}
	}
}

//根据索引号数组得到元素数组
func findNumbsByIndexes(nums []int, indexes [][]int) [][]int {
	if len(indexes) == 0 {
		return [][]int{}
	}
	result := make([][]int, len(indexes))
	for i, v := range indexes {
		line := make([]int, 0)
		for j, v2 := range v {
			if v2 == 1 {
				line = append(line, nums[j])
			}
		}
		result[i] = line
	}
	return result
}

//数学方法计算组合数(从n中取m个数)
func mathCombination(n int, m int) int {
	return factorial(n) / (factorial(n-m) * factorial(m))
}

//阶乘
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
