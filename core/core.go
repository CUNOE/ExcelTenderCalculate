package core

type Result struct {
	Combination []int
	AvgPrice    float64
	Companies   []InRandomPriceWinTheBidding
	Count       map[string]int
}

type InRandomPriceWinTheBidding struct {
	RandomPrice             float64
	TheCompanyWinTheBidding Company
}

func CalculateResult(combination [][]int, select_companys []Company, fixed_companys []Company, randomValue []float64) []Result {

	var r []Result

	for _, z := range combination {
		var prices []float64
		var cs []Company
		var zhongbiaocompany []InRandomPriceWinTheBidding
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
			zhongbiaocompany = append(zhongbiaocompany, InRandomPriceWinTheBidding{
				RandomPrice:             rp,
				TheCompanyWinTheBidding: ReturnCompanyWhoWinTheBidding(rp, cs),
			})
		}
		r = append(r, Result{
			Combination: z,
			AvgPrice:    AvgPrice(prices),
			Companies:   zhongbiaocompany,
			Count:       CountTheResult(zhongbiaocompany),
		})
	}

	return r
}

func CountTheResult(zbc []InRandomPriceWinTheBidding) map[string]int {
	var count map[string]int
	count = make(map[string]int)

	for _, d := range zbc {
		name := d.TheCompanyWinTheBidding.Name
		if _, ok := count[name]; ok {
			count[name]++
		} else {
			count[name] = 1
		}
	}
	return count
}
