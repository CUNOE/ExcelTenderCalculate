package core

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

// Output 将数据写入Excel文件
func Output(results []Result, selectCompanies []Company, fixedCompanies []Company, randomValue []float64, path string) error {
	f := excelize.NewFile()

	r := 3
	s := 3
	d := 3

	g := 2

	start := 3
	end := 3

	f.SetPanes("Sheet1", `{
		"freeze":true,
		"split":false,
		"x_split":0,
		"y_split":1,
		"top_left_cell": "A2"
	}`)

	styleMid, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#D9D9D9"}, Pattern: 1},
	})

	styleTitle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Size: 12},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		//Fill: excelize.Fill{Type: "pattern", Color: []string{"#FFFF00"}, Pattern: 1},
	})

	styleBody, _ := f.NewStyle(&excelize.Style{
		//Fill: excelize.Fill{Color: []string{"#F2F2F2"}, Pattern: 1},
		//Font: &excelize.Font{
		//	Bold: true,
		//	//Color: "#FA7D00",
		//},
		//Fill: excelize.Fill{Type: "pattern", Color: []string{"#F0F0F0"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})

	//styleFill, _ := f.NewStyle(&excelize.Style{
	//	Fill: excelize.Fill{Type: "pattern", Color: []string{"#E0EBF5"}, Pattern: 1},
	//})

	//styleBorder, _ := f.NewStyle(&excelize.Style{
	//	Border: []excelize.Border{
	//		{Type: "left", Color: "#000000", Style: 2},
	//		{Type: "right", Color: "#000000", Style: 2},
	//		{Type: "top", Color: "#000000", Style: 2},
	//		{Type: "bottom", Color: "#000000", Style: 2},
	//	},
	//})

	f.SetCellStyle("Sheet1", "A1", "K1", styleTitle)

	f.SetColWidth("Sheet1", "B", "B", 25)
	f.SetColWidth("Sheet1", "C", "C", 10)
	f.SetColWidth("Sheet1", "D", "D", 10)
	f.SetColWidth("Sheet1", "E", "E", 12)
	f.SetColWidth("Sheet1", "F", "F", 17)
	f.SetColWidth("Sheet1", "G", "G", 10)
	f.SetColWidth("Sheet1", "H", "H", 10)
	f.SetColWidth("Sheet1", "I", "I", 10)
	f.SetColWidth("Sheet1", "J", "J", 10)
	f.SetColWidth("Sheet1", "K", "K", 15)

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
			end++
		}
		//log.Println(end)

		//f.SetCellStyle("Sheet1", "A"+strconv.Itoa(start), "K"+strconv.Itoa(end-1), styleBorder)
		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(start), "J"+strconv.Itoa(end-1), styleBody)
		//f.SetCellStyle("Sheet1", "A"+strconv.Itoa(start), "K"+strconv.Itoa(end-1), styleFill)

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(g), "J"+strconv.Itoa(g), styleMid)
		g = s
		if len(z.Companies) >= len(z.Combination)+len(fixedCompanies) {
			s++
			r = s
			d = s
			start = s
			end = s
		} else {
			r++
			s = r
			d = r
			start = r
			end = r
		}

	}
	err := f.SaveAs(path)
	if err != nil {
		log.Printf("请检查excel是否完全关闭output.xlsx文件")
		return err
	}
	return nil
}
