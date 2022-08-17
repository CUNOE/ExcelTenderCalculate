package core

import (
	"github.com/xuri/excelize/v2"
	"log"
)

// Output 将数据写入Excel文件
func Output(results []Result, selectCompanies []Company, fixedCompanies []Company, randomValue []float64, path string) error {
	f := excelize.NewFile()

	r := 2
	s := 2
	d := 2

	//f.SetPanes("Sheet1", `{
	//	"freeze":true,
	//	"split":false,
	//	"x_split":0,
	//	"y_split":1,
	//	"top_left_cell": "A2",
	//}`)

	f.SetColWidth("Sheet1", "B", "B", 25)
	f.SetColWidth("Sheet1", "E", "E", 10)
	f.SetColWidth("Sheet1", "F", "F", 15)

	cell, _ := excelize.CoordinatesToCellName(1, 1)
	f.SetCellValue("Sheet1", cell, "编号")
	cell, _ = excelize.CoordinatesToCellName(2, 1)
	f.SetCellValue("Sheet1", cell, "公司名称")
	cell, _ = excelize.CoordinatesToCellName(3, 1)
	f.SetCellValue("Sheet1", cell, "投标价格")
	cell, _ = excelize.CoordinatesToCellName(4, 1)
	f.SetCellValue("Sheet1", cell, "平均价格")
	cell, _ = excelize.CoordinatesToCellName(5, 1)
	f.SetCellValue("Sheet1", cell, "随机费率%")
	cell, _ = excelize.CoordinatesToCellName(6, 1)
	f.SetCellValue("Sheet1", cell, "随机相对应报价")
	cell, _ = excelize.CoordinatesToCellName(7, 1)
	f.SetCellValue("Sheet1", cell, "中标公司")
	cell, _ = excelize.CoordinatesToCellName(8, 1)
	f.SetCellValue("Sheet1", cell, "中标价格")
	cell, _ = excelize.CoordinatesToCellName(9, 1)
	f.SetCellValue("Sheet1", cell, "中标公司")
	cell, _ = excelize.CoordinatesToCellName(10, 1)
	f.SetCellValue("Sheet1", cell, "中标次数")
	cell, _ = excelize.CoordinatesToCellName(11, 1)
	f.SetCellValue("Sheet1", cell, "本次共有组合数")

	cell, _ = excelize.CoordinatesToCellName(11, 2)
	f.SetCellValue("Sheet1", cell, len(results))

	for _, z := range results {
		// 平均价铬
		cell, _ := excelize.CoordinatesToCellName(4, r)
		f.SetCellValue("Sheet1", cell, z.AvgPrice)

		// 中标次数统计
		for k, v := range z.Count {
			cell, _ = excelize.CoordinatesToCellName(9, d)
			f.SetCellValue("Sheet1", cell, k)
			cell, _ = excelize.CoordinatesToCellName(10, d)
			f.SetCellValue("Sheet1", cell, v)
			d++
		}

		// 投标公司统计
		for _, fc := range fixedCompanies {
			cell, _ := excelize.CoordinatesToCellName(1, r)
			f.SetCellValue("Sheet1", cell, fc.ID)
			cell, _ = excelize.CoordinatesToCellName(2, r)
			f.SetCellValue("Sheet1", cell, fc.Name)
			cell, _ = excelize.CoordinatesToCellName(3, r)
			f.SetCellValue("Sheet1", cell, fc.Price)
			r++
		}

		// 投标公司统计
		for _, zh := range z.Combination {
			cell, _ := excelize.CoordinatesToCellName(1, r)
			f.SetCellValue("Sheet1", cell, selectCompanies[zh].ID)
			cell, _ = excelize.CoordinatesToCellName(2, r)
			f.SetCellValue("Sheet1", cell, selectCompanies[zh].Name)
			cell, _ = excelize.CoordinatesToCellName(3, r)
			f.SetCellValue("Sheet1", cell, selectCompanies[zh].Price)
			r++
		}

		// 随机费率及相对应报价 && 对应中标公司
		for n, zbc := range z.Companies {
			cell, _ := excelize.CoordinatesToCellName(5, s)
			f.SetCellValue("Sheet1", cell, randomValue[n])
			cell, _ = excelize.CoordinatesToCellName(6, s)
			f.SetCellValue("Sheet1", cell, zbc.RandomPrice)
			cell, _ = excelize.CoordinatesToCellName(7, s)
			f.SetCellValue("Sheet1", cell, zbc.TheCompanyWinTheBidding.Name)
			cell, _ = excelize.CoordinatesToCellName(8, s)
			f.SetCellValue("Sheet1", cell, zbc.TheCompanyWinTheBidding.Price)
			s++
		}
		if len(z.Companies) >= len(z.Combination)+len(fixedCompanies) {
			s++
			r = s
			d = s
		} else {
			r++
			s = r
			d = r
		}

	}
	err := f.SaveAs(path)
	if err != nil {
		log.Printf("请检查excel是否完全关闭output.xlsx文件")
		return err
	}
	return nil
}
