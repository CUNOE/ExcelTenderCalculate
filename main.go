package main

import (
	"Biao/core"
	"fmt"
	"log"
	"time"
)

func main() {

	start := time.Now()

	// 从excel表格中读取投标单位数据以及随机费率数据
	log.Println("读取输入数据...")
	companies, randomValue, fixed, m := core.ReadExcel("input.xlsx")
	log.Printf("读取完成，共读取%d条数据\n", len(companies))
	log.Printf("投标单位%v\n", companies)
	log.Printf("随机费率数据：%v\n", randomValue)
	log.Printf("固定的投标单位编号：%v\n", fixed)

	var select_companys []core.Company
	var fixed_companys []core.Company

	for _, v := range companies {
		if core.IsContain(fixed, v.ID) {
			fixed_companys = append(fixed_companys, v)
		} else {
			select_companys = append(select_companys, v)
		}

	}
	log.Printf("可选择的投标单位%v\n", select_companys)
	log.Printf("固定的投标单位%v\n", fixed_companys)

	combination := core.Combination(core.GenerateNums(len(select_companys)), m-len(fixed_companys))
	log.Printf("组合数量：%v\n", len(combination))

	result := core.CalculateResult(combination, select_companys, fixed_companys, randomValue)

	log.Printf("计算完成，共计算%d条数据\n", len(result))
	log.Printf("正在写入数据...\n")
	log.Printf("共需要写入%d条数据\n", len(result)*len(randomValue)+len(result))
	err := core.Output(result, select_companys, fixed_companys, randomValue, "output.xlsx")
	if err != nil {
		log.Printf("写入失败，错误信息：%v\n", err)
		fmt.Println("按任意键退出...")
		var input string
		fmt.Scanln(&input)
	}
	log.Printf("写入完成\n")

	log.Printf("任务完成，耗时%s\n", time.Since(start))
	fmt.Println("按任意键退出...")
	var input string
	fmt.Scanln(&input)

}
