package core

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"sort"
	"strconv"
)

type Company struct {
	ID    int
	Name  string
	Price float64
}

// 读取Excel文件
func ReadExcel(path string) ([]Company, []float64, []int, int) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return nil, nil, nil, 0
	}

	var randomValue []float64
	var companys []Company
	var fixed []int
	var m int

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Printf("请检查input.xlsx文件是否存在Sheet1表格\n")
		log.Fatal("读取Excel文件失败，错误信息：%v\n", err)
		return nil, nil, nil, 0
	}

	for d, row := range rows {
		// 读取需要抽取的总数
		if d == 1 {
			for k, r := range row {
				if k == 7 {
					l, _ := strconv.ParseInt(r, 10, 64)
					m = int(l)
				}
			}
		}
		if d >= 1 {
			for k, r := range row {
				// 读取随机费率
				if k == 4 {
					v, _ := strconv.ParseFloat(r, 64)
					randomValue = append(randomValue, v)
				}
				// 读取投标公司
				if k == 2 {
					p, _ := strconv.ParseFloat(r, 64)
					companys = append(companys, Company{
						ID:    d,
						Name:  row[1],
						Price: p,
					})
				}

				// 读取固定公司
				if k == 6 {
					f, _ := strconv.ParseInt(r, 10, 64)
					fixed = append(fixed, int(f))
				}

			}

		}

	}
	fixed = DeleteSlice(fixed)
	fixed = RemoveDuplicatesInPlace(fixed)

	return companys, randomValue, fixed, m

}

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
