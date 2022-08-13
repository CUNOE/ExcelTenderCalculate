package core

import (
	"github.com/xuri/excelize/v2"
	"log"
)

//func Save(zbzh []ZhongBiaoZuHe, companys []Company, path string, randomValue []float64) error {
//	f := excelize.NewFile()
//
//	r := 2
//	s := 2
//	d := 2
//
//	//f.SetPanes("Sheet1", `{
//	//	"freeze":true,
//	//	"split":false,
//	//	"x_split":0,
//	//	"y_split":1,
//	//	"top_left_cell": "A2",
//	//}`)
//
//	cell, _ := excelize.CoordinatesToCellName(1, 1)
//	f.SetCellValue("Sheet1", cell, "招标公司编号")
//	cell, _ = excelize.CoordinatesToCellName(2, 1)
//	f.SetCellValue("Sheet1", cell, "招标公司名称")
//	cell, _ = excelize.CoordinatesToCellName(3, 1)
//	f.SetCellValue("Sheet1", cell, "招标公司价格")
//	cell, _ = excelize.CoordinatesToCellName(4, 1)
//	f.SetCellValue("Sheet1", cell, "平均价格")
//	cell, _ = excelize.CoordinatesToCellName(5, 1)
//	f.SetCellValue("Sheet1", cell, "随机费率%")
//	cell, _ = excelize.CoordinatesToCellName(6, 1)
//	f.SetCellValue("Sheet1", cell, "随机相对应报价")
//	cell, _ = excelize.CoordinatesToCellName(7, 1)
//	f.SetCellValue("Sheet1", cell, "中标公司")
//	cell, _ = excelize.CoordinatesToCellName(8, 1)
//	f.SetCellValue("Sheet1", cell, "中标价格")
//	cell, _ = excelize.CoordinatesToCellName(9, 1)
//	f.SetCellValue("Sheet1", cell, "中标公司")
//	cell, _ = excelize.CoordinatesToCellName(10, 1)
//	f.SetCellValue("Sheet1", cell, "中标次数")
//	cell, _ = excelize.CoordinatesToCellName(11, 1)
//	f.SetCellValue("Sheet1", cell, "本次共有组合数")
//	cell, _ = excelize.CoordinatesToCellName(11, 2)
//	f.SetCellValue("Sheet1", cell, len(zbzh))
//
//	for _, z := range zbzh {
//		// 平均价铬
//		cell, _ := excelize.CoordinatesToCellName(4, r)
//		f.SetCellValue("Sheet1", cell, z.Avgprice)
//
//		// 中标次数统计
//		for k, v := range z.Count {
//			cell, _ = excelize.CoordinatesToCellName(9, d)
//			f.SetCellValue("Sheet1", cell, k)
//			cell, _ = excelize.CoordinatesToCellName(10, d)
//			f.SetCellValue("Sheet1", cell, v)
//			d++
//		}
//
//		// 投标公司统计
//		for _, zh := range z.Zuhe {
//			cell, _ := excelize.CoordinatesToCellName(1, r)
//			f.SetCellValue("Sheet1", cell, companys[zh].ID)
//			cell, _ = excelize.CoordinatesToCellName(2, r)
//			f.SetCellValue("Sheet1", cell, companys[zh].Name)
//			cell, _ = excelize.CoordinatesToCellName(3, r)
//			f.SetCellValue("Sheet1", cell, companys[zh].Price)
//			r++
//		}
//
//		// 中标公司统计
//		for n, zbc := range z.ZhongbiaoCompanys {
//			cell, _ := excelize.CoordinatesToCellName(5, s)
//			f.SetCellValue("Sheet1", cell, randomValue[n])
//			cell, _ = excelize.CoordinatesToCellName(6, s)
//			f.SetCellValue("Sheet1", cell, zbc.Randomprice)
//			cell, _ = excelize.CoordinatesToCellName(7, s)
//			f.SetCellValue("Sheet1", cell, zbc.Zhongbiao.Name)
//			cell, _ = excelize.CoordinatesToCellName(8, s)
//			f.SetCellValue("Sheet1", cell, zbc.Zhongbiao.Price)
//			s++
//		}
//		if len(z.ZhongbiaoCompanys) >= len(z.Zuhe) {
//			s++
//			r = s
//			d = s
//		} else {
//			r++
//			s = r
//			d = r
//		}
//
//	}
//	err := f.SaveAs(path)
//	if err != nil {
//		log.Printf("请检查excel是否完全关闭output.xlsx文件")
//		return err
//	}
//	return nil
//}

func Output(zbzh []ZhongBiaoZuHe, select_companys []Company, fixed_companys []Company, randomValue []float64, path string) error {
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

	cell, _ := excelize.CoordinatesToCellName(1, 1)
	f.SetCellValue("Sheet1", cell, "招标公司编号")
	cell, _ = excelize.CoordinatesToCellName(2, 1)
	f.SetCellValue("Sheet1", cell, "招标公司名称")
	cell, _ = excelize.CoordinatesToCellName(3, 1)
	f.SetCellValue("Sheet1", cell, "招标公司价格")
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
	f.SetCellValue("Sheet1", cell, len(zbzh))

	for _, z := range zbzh {
		// 平均价铬
		cell, _ := excelize.CoordinatesToCellName(4, r)
		f.SetCellValue("Sheet1", cell, z.Avgprice)

		// 中标次数统计
		for k, v := range z.Count {
			cell, _ = excelize.CoordinatesToCellName(9, d)
			f.SetCellValue("Sheet1", cell, k)
			cell, _ = excelize.CoordinatesToCellName(10, d)
			f.SetCellValue("Sheet1", cell, v)
			d++
		}

		// 投标公司统计
		for _, fc := range fixed_companys {
			cell, _ := excelize.CoordinatesToCellName(1, r)
			f.SetCellValue("Sheet1", cell, fc.ID)
			cell, _ = excelize.CoordinatesToCellName(2, r)
			f.SetCellValue("Sheet1", cell, fc.Name)
			cell, _ = excelize.CoordinatesToCellName(3, r)
			f.SetCellValue("Sheet1", cell, fc.Price)
			r++
		}

		// 投标公司统计
		for _, zh := range z.Zuhe {
			cell, _ := excelize.CoordinatesToCellName(1, r)
			f.SetCellValue("Sheet1", cell, select_companys[zh].ID)
			cell, _ = excelize.CoordinatesToCellName(2, r)
			f.SetCellValue("Sheet1", cell, select_companys[zh].Name)
			cell, _ = excelize.CoordinatesToCellName(3, r)
			f.SetCellValue("Sheet1", cell, select_companys[zh].Price)
			r++
		}

		// 随机费率及相对应报价 && 对应中标公司
		for n, zbc := range z.ZhongbiaoCompanys {
			cell, _ := excelize.CoordinatesToCellName(5, s)
			f.SetCellValue("Sheet1", cell, randomValue[n])
			cell, _ = excelize.CoordinatesToCellName(6, s)
			f.SetCellValue("Sheet1", cell, zbc.Randomprice)
			cell, _ = excelize.CoordinatesToCellName(7, s)
			f.SetCellValue("Sheet1", cell, zbc.Zhongbiao.Name)
			cell, _ = excelize.CoordinatesToCellName(8, s)
			f.SetCellValue("Sheet1", cell, zbc.Zhongbiao.Price)
			s++
		}
		if len(z.ZhongbiaoCompanys) >= len(z.Zuhe)+len(fixed_companys) {
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
