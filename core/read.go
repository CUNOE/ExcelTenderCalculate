package core

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

type Company struct {
	ID      int
	Name    string
	Price   float64
	IsFixed bool
}

// ReadExcel 读取Excel文件
func ReadExcel(path string) ([]Company, []float64, int) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return nil, nil, 0
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var randomValue []float64
	var companies []Company
	var m int

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Printf("请检查input.xlsx文件是否存在Sheet1表格\n")
		log.Fatalf("读取Excel文件失败，错误信息：%v\n", err)
		return nil, nil, 0
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
				if k == 6 && r != "" {
					v, _ := strconv.ParseFloat(r, 64)
					randomValue = append(randomValue, v)
				}
			}

			// 读取公司信息
			if row[0] == "TRUE" && row[1] != "" && row[2] != "" && row[3] != "" {
				if row[4] == "TRUE" {
					id, _ := strconv.ParseInt(row[1], 10, 64)
					price, _ := strconv.ParseFloat(row[3], 64)
					companies = append(companies, Company{int(id), row[2], price, true})
				} else {
					id, _ := strconv.ParseInt(row[1], 10, 64)
					price, _ := strconv.ParseFloat(row[3], 64)
					companies = append(companies, Company{int(id), row[2], price, false})
				}
			}

		}

	}

	return companies, randomValue, m

}
