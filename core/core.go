package core

type ZhongBiaoZuHe struct {
	Zuhe              []int
	Avgprice          float64
	ZhongbiaoCompanys []ZhongbiaoCompany
	Count             map[string]int
}

type ZhongbiaoCompany struct {
	Randomprice float64
	Zhongbiao   Company
}

func CalculateZhongbiao(zuhe [][]int, companys []Company, randomValue []float64) []ZhongBiaoZuHe {

	var zhongbiao []ZhongBiaoZuHe

	for _, z := range zuhe {
		var prices []float64
		var cs []Company
		var zhongbiaocompany []ZhongbiaoCompany
		for _, c := range z {
			prices = append(prices, companys[c].Price)
			cs = append(cs, companys[c])
		}
		randomPrices := CalculateRandomPrices(AvgPrice(prices), randomValue)

		for _, rp := range randomPrices {
			zhongbiaocompany = append(zhongbiaocompany, ZhongbiaoCompany{
				Randomprice: rp,
				Zhongbiao:   Zhongbiao(rp, cs),
			})
		}
		zhongbiao = append(zhongbiao, ZhongBiaoZuHe{
			Zuhe:              z,
			Avgprice:          AvgPrice(prices),
			ZhongbiaoCompanys: zhongbiaocompany,
			Count:             CountZhongbiao(zhongbiaocompany),
		})
	}

	return zhongbiao
}

func CalculateZhongbiaoWithFixed(zuhe [][]int, select_companys []Company, fixed_companys []Company, randomValue []float64) []ZhongBiaoZuHe {

	var zhongbiao []ZhongBiaoZuHe

	for _, z := range zuhe {
		var prices []float64
		var cs []Company
		var zhongbiaocompany []ZhongbiaoCompany
		for _, c := range z {
			prices = append(prices, select_companys[c].Price)
			cs = append(cs, select_companys[c])
		}
		for _, c := range fixed_companys {
			prices = append(prices, c.Price)
			cs = append(cs, c)
		}
		randomPrices := CalculateRandomPrices(AvgPrice(prices), randomValue)

		for _, rp := range randomPrices {
			zhongbiaocompany = append(zhongbiaocompany, ZhongbiaoCompany{
				Randomprice: rp,
				Zhongbiao:   Zhongbiao(rp, cs),
			})
		}
		zhongbiao = append(zhongbiao, ZhongBiaoZuHe{
			Zuhe:              z,
			Avgprice:          AvgPrice(prices),
			ZhongbiaoCompanys: zhongbiaocompany,
			Count:             CountZhongbiao(zhongbiaocompany),
		})
	}

	return zhongbiao
}

func CountZhongbiao(zbc []ZhongbiaoCompany) map[string]int {
	var count map[string]int
	count = make(map[string]int)

	for _, d := range zbc {
		name := d.Zhongbiao.Name
		if _, ok := count[name]; ok {
			count[name]++
		} else {
			count[name] = 1
		}
	}
	return count
}
